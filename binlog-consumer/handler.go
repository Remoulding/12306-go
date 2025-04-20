package main

import (
	"context"
	"fmt"
	"github.com/Remoulding/12306-go/ticket-service/biz/dao"
	"github.com/Remoulding/12306-go/ticket-service/biz/service"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/go-mysql-org/go-mysql/schema"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"strings"
	"sync"
)

// SeatStatus 座位状态常量
const (
	SeatAvailable = 0 // 座位可用
	SeatLocked    = 1 // 座位已锁定
)

// SeatEventHandler 处理座位变更事件
type SeatEventHandler struct {
	redisClient *redis.Client
}

// NewSeatEventHandler 创建新的座位事件处理器
func NewSeatEventHandler(redisClient *redis.Client) *SeatEventHandler {
	return &SeatEventHandler{
		redisClient: redisClient,
	}
}

func (h *SeatEventHandler) OnPosSynced(header *replication.EventHeader, pos mysql.Position, _ mysql.GTIDSet, _ bool) error {
	ctx := context.Background()
	h.redisClient.Set(ctx, "canal:position", fmt.Sprintf("%s|%d", pos.Name, pos.Pos), 0)
	return nil
}

// LoadPosition 启动时加载历史位置
func (h *SeatEventHandler) LoadPosition() mysql.Position {
	val, err := h.redisClient.Get(context.Background(), "canal:position").Result()
	if err != nil {
		return mysql.Position{} // 默认从头开始
	}

	parts := strings.Split(val, "|")
	pos, _ := strconv.ParseUint(parts[1], 10, 32)
	return mysql.Position{Name: parts[0], Pos: uint32(pos)}
}

// OnRow 处理行变更事件
func (h *SeatEventHandler) OnRow(e *canal.RowsEvent) error {
	log.Printf("Action: %s, Table: %s.%s", e.Action, e.Table.Schema, e.Table.Name)
	// 只处理UPDATE事件
	if e.Action != canal.UpdateAction {
		return nil
	}

	// 确保处理的是t_seat表
	if e.Table.Name != "t_seat" {
		return nil
	}

	// 解析变更前后的数据
	wg := &sync.WaitGroup{}
	for i := 0; i < len(e.Rows); i += 2 {
		oldRow := e.Rows[i]
		newRow := e.Rows[i+1]

		// 将行数据转为map
		oldData := h.rowToMap(e.Table, oldRow)
		newData := h.rowToMap(e.Table, newRow)

		log.Printf("oldData: %v, newData: %v\n", oldData, newData)

		// 检查SeatStatus是否变化
		oldStatus, oldOk := oldData["seat_status"].(int32)
		newStatus, newOk := newData["seat_status"].(int32)
		if !oldOk || !newOk || oldStatus == newStatus {
			continue
		}

		// 判断状态流转方向
		switch {
		case oldStatus == SeatAvailable && newStatus == SeatLocked:
			// 锁定座位：扣减库存
			go func() {
				wg.Add(1)
				defer wg.Done()
				h.updateRedisStock(newData, -1)
			}()
		case oldStatus == SeatLocked && newStatus == SeatAvailable:
			// 释放座位：增加库存
			go func() {
				wg.Add(1)
				defer wg.Done()
				h.updateRedisStock(newData, 1)
			}()
		}
	}
	wg.Wait()
	return nil
}

// rowToMap 将行数据转为map
func (h *SeatEventHandler) rowToMap(table *schema.Table, row []interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i, col := range table.Columns {
		result[col.Name] = row[i]
	}
	return result
}

// updateRedisStock 更新Redis库存
func (h *SeatEventHandler) updateRedisStock(data map[string]interface{}, delta int) {
	// 构造Redis Key (trainID_StartStation_EndStation_DepartureDate)
	tid := data["train_id"].(int64)
	trainID := strconv.Itoa(int(tid))
	startStation := data["start_station"].(string)
	endStation := data["end_station"].(string)
	departureDate := data["departure_date"].(string)
	key := fmt.Sprintf(configs.TrainRemainTicket,
		trainID,
		startStation,
		endStation,
		departureDate,
	)

	// 获取座位类型
	seatType, ok := data["seat_type"].(int32)
	if !ok {
		log.Printf("无效的seat_type: %v", data["seat_type"])
		return
	}

	// 原子性更新库存
	ctx := context.Background()
	exists, err2 := h.redisClient.Exists(ctx, key).Result()
	if err2 != nil {
		configs.Log.WithContext(ctx).Errorf("redis.Exists failed, err: %v", err2)
		return
	}
	if exists == 0 {
		log.Printf("Redis key不存在: %s", key)
		lockKey := fmt.Sprintf(configs.LockTrainRemainTicket,
			trainID,
			startStation,
			endStation,
			departureDate,
		)
		_, err := service.SafeLoadHash(ctx, key, lockKey, func(ctx context.Context) (map[string]string, error) {
			groupedByType, err3 := dao.QuerySeatCountGroupedByType(ctx, tid, startStation, endStation, departureDate)
			if err3 != nil {
				return nil, err3
			}
			mp := make(map[string]string)
			for _, seatCount := range groupedByType {
				mp[strconv.Itoa(seatCount.SeatType)] = strconv.Itoa(int(seatCount.Count))
			}
			return mp, nil
		})
		if err != nil {
			configs.Log.WithContext(ctx).Errorf("query seat count failed, err: %v", err)
		}
		return
	}
	err := h.redisClient.HIncrBy(ctx, key, strconv.Itoa(int(seatType)), int64(delta)).Err()
	if err != nil {
		log.Printf("Redis更新失败: key=%s, seatType=%d, error=%v", key, seatType, err)
		// 这里可以加入重试或报警逻辑
	} else {
		log.Printf("库存更新成功: key=%s, seatType=%d, delta=%d", key, seatType, delta)
	}
	return
}

// String 返回处理器名称
func (h *SeatEventHandler) String() string {
	return "SeatEventHandler"
}

func (h *SeatEventHandler) OnRotate(*replication.EventHeader, *replication.RotateEvent) error {
	return nil
}
func (h *SeatEventHandler) OnTableChanged(*replication.EventHeader, string, string) error { return nil }
func (h *SeatEventHandler) OnDDL(*replication.EventHeader, mysql.Position, *replication.QueryEvent) error {
	return nil
}
func (h *SeatEventHandler) OnXID(*replication.EventHeader, mysql.Position) error         { return nil }
func (h *SeatEventHandler) OnGTID(*replication.EventHeader, mysql.BinlogGTIDEvent) error { return nil }
func (h *SeatEventHandler) OnRowsQueryEvent(*replication.RowsQueryEvent) error           { return nil }

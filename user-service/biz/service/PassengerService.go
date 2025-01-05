package service

import (
	"context"
	"encoding/json"
	"fmt"
	. "github.com/Remoulding/12306-go/user-service/biz/dto/req"
	. "github.com/Remoulding/12306-go/user-service/biz/dto/resp"
	. "github.com/Remoulding/12306-go/user-service/biz/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// PassengerService 乘车人服务接口
type PassengerService interface {

	// ListPassengerQueryByUsername 根据用户名查询乘车人列表
	ListPassengerQueryByUsername(username string) ([]PassengerRespDTO, error)

	// ListPassengerQueryByIds 根据乘车人 ID 集合查询乘车人列表
	ListPassengerQueryByIds(username string, ids []int64) ([]PassengerActualRespDTO, error)

	// SavePassenger 新增乘车人
	SavePassenger(requestParam PassengerReqDTO) error

	// UpdatePassenger 修改乘车人
	UpdatePassenger(requestParam PassengerReqDTO) error

	// RemovePassenger 移除乘车人
	RemovePassenger(requestParam PassengerRemoveReqDTO) error
}

// PassengerServiceImpl 乘车人服务实现
type PassengerServiceImpl struct {
	ctx            context.Context
	DB             *gorm.DB
	Cache          *redis.Client
	CacheTTL       time.Duration
	TransactionMgr *gorm.DB
}

// NewPassengerServiceImpl 创建乘车人服务实现
func NewPassengerServiceImpl(ctx context.Context, db *gorm.DB, cache *redis.Client, cacheTTL time.Duration, transactionMgr *gorm.DB) *PassengerServiceImpl {
	return &PassengerServiceImpl{
		ctx:            ctx,
		DB:             db,
		Cache:          cache,
		CacheTTL:       cacheTTL,
		TransactionMgr: transactionMgr,
	}
}

// 常量定义
const CacheKeyPassengerList = "USER_PASSENGER_LIST"

// ListPassengerQueryByUsername 根据用户名查询乘车人列表
func (s *PassengerServiceImpl) ListPassengerQueryByUsername(username string) ([]PassengerRespDTO, error) {
	cacheKey := fmt.Sprintf("%s_%s", CacheKeyPassengerList, username)
	cachedData, err := s.Cache.Get(s.ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var passengers []PassengerRespDTO
		err = json.Unmarshal([]byte(cachedData), &passengers)
		if err == nil {
			return passengers, nil
		}
	}

	var passengers []PassengerDO
	if err := s.DB.Where("username = ?", username).Find(&passengers).Error; err != nil {
		return nil, err
	}

	resp := make([]PassengerRespDTO, len(passengers))
	for i, passenger := range passengers {
		resp[i] = PassengerRespDTO{
			ID:       strconv.Itoa(int(passenger.ID)),
			RealName: passenger.RealName,
		}
	}

	data, _ := json.Marshal(resp)
	s.Cache.Set(s.ctx, cacheKey, data, s.CacheTTL)
	return resp, nil
}

// ListPassengerQueryByIds 根据乘车人 ID 集合查询乘车人列表
func (s *PassengerServiceImpl) ListPassengerQueryByIds(username string, ids []int64) ([]PassengerActualRespDTO, error) {
	passengers, err := s.ListPassengerQueryByUsername(username)
	if err != nil {
		return nil, err
	}

	var filtered []PassengerActualRespDTO
	for _, passenger := range passengers {
		for _, id := range ids {
			if passenger.ID == strconv.Itoa(int(id)) {
				filtered = append(filtered, PassengerActualRespDTO{
					ID:       passenger.ID,
					RealName: passenger.RealName,
				})
				break
			}
		}
	}
	return filtered, nil
}

// SavePassenger 新增乘车人
func (s *PassengerServiceImpl) SavePassenger(requestParam PassengerReqDTO) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		id, err := strconv.ParseInt(requestParam.ID, 10, 64)
		if err != nil {
			// TODO
		}
		passenger := PassengerDO{
			ID:           id,
			RealName:     requestParam.RealName,
			Username:     "currentUsername", // Replace with actual username logic
			CreateDate:   time.Now(),
			VerifyStatus: 1,
		}
		if err := tx.Create(&passenger).Error; err != nil {
			return err
		}
		return s.invalidateCache("currentUsername") // Replace with actual username logic
	})
}

// UpdatePassenger 修改乘车人
func (s *PassengerServiceImpl) UpdatePassenger(requestParam PassengerReqDTO) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&PassengerDO{}).Where("id = ?", requestParam.ID).Updates(map[string]interface{}{
			"real_name": requestParam.RealName,
		}).Error; err != nil {
			return err
		}
		return s.invalidateCache("currentUsername") // Replace with actual username logic
	})
}

// RemovePassenger 移除乘车人
func (s *PassengerServiceImpl) RemovePassenger(requestParam PassengerRemoveReqDTO) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&PassengerDO{}).Where("id = ?", requestParam.ID).Update("del_flag", true).Error; err != nil {
			return err
		}
		return s.invalidateCache("currentUsername") // Replace with actual username logic
	})
}

// Helper 方法
func (s *PassengerServiceImpl) invalidateCache(username string) error {
	cacheKey := fmt.Sprintf("%s_%s", CacheKeyPassengerList, username)
	return s.Cache.Del(s.ctx, cacheKey).Err()
}

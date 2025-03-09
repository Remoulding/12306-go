package model

import "time"

// TrainStationRelationDO 列车车站关联实体
type TrainStationRelationDO struct {
	BaseDO                  // 嵌入 BaseDO 模拟继承
	ID            int64     `gorm:"primaryKey;column:id"`  // id
	TrainID       int64     `gorm:"column:train_id"`       // 车次id
	Departure     string    `gorm:"column:departure"`      // 出发站点
	Arrival       string    `gorm:"column:arrival"`        // 到达站点
	StartRegion   string    `gorm:"column:start_region"`   // 起始城市
	EndRegion     string    `gorm:"column:end_region"`     // 终点城市
	DepartureFlag bool      `gorm:"column:departure_flag"` // 始发站标识
	ArrivalFlag   bool      `gorm:"column:arrival_flag"`   // 终点站标识
	DepartureTime time.Time `gorm:"column:departure_time"` // 出发时间
	ArrivalTime   time.Time `gorm:"column:arrival_time"`   // 到达时间
}

// TableName 设置表名
func (TrainStationRelationDO) TableName() string {
	return "t_train_station_relation"
}

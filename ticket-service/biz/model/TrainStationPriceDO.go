package model

// TrainStationPriceDO 列车车站票价实体
type TrainStationPriceDO struct {
	BaseDO           // 嵌入 BaseDO 模拟继承
	ID        int64  `gorm:"primaryKey;column:id"` // id
	TrainID   int64  `gorm:"column:train_id"`      // 车次id
	SeatType  int    `gorm:"column:seat_type"`     // 座位类型
	Departure string `gorm:"column:departure"`     // 出发站点
	Arrival   string `gorm:"column:arrival"`       // 到达站点
	Price     int    `gorm:"column:price"`         // 车票价格
}

// TableName 设置表名
func (TrainStationPriceDO) TableName() string {
	return "t_train_station_price"
}

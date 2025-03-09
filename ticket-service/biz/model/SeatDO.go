package model

// SeatDO 座位实体
type SeatDO struct {
	BaseDO                // 嵌入 BaseDO 模拟继承
	ID             int64  `gorm:"primaryKey;column:id"`   // id
	TrainID        int64  `gorm:"column:train_id"`        // 列车id
	CarriageNumber string `gorm:"column:carriage_number"` // 车厢号
	SeatNumber     string `gorm:"column:seat_number"`     // 座位号
	SeatType       int    `gorm:"column:seat_type"`       // 座位类型， 0:商务；1:一等；2:二等
	StartStation   string `gorm:"column:start_station"`   // 起始站
	EndStation     string `gorm:"column:end_station"`     // 终点站
	SeatStatus     int    `gorm:"column:seat_status"`     // 座位状态
	Price          int    `gorm:"column:price"`           // 车票价格
}

// TableName 设置表名
func (SeatDO) TableName() string {
	return "t_seat"
}

const (
	SeatAvailable = 0 // 座位可用
	SeatLocked    = 1 // 座位已锁定
	SeatSold      = 2 // 座位已售出
)

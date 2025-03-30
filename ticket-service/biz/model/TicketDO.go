package model

import "time"

// TicketDO 车票实体
type TicketDO struct {
	BaseDO                   // 嵌入 BaseDO 模拟继承
	ID             int64     `gorm:"primaryKey;column:id"`   // id
	Username       string    `gorm:"column:username"`        // 用户名
	TrainID        int64     `gorm:"column:train_id"`        // 列车id
	TrainNumber    string    `gorm:"column:train_number"`    // 列车车次
	CarriageNumber string    `gorm:"column:carriage_number"` // 车厢号
	SeatNumber     string    `gorm:"column:seat_number"`     // 座位号
	PassengerID    int64     `gorm:"column:passenger_id"`    // 乘车人 ID
	TicketStatus   int       `gorm:"column:ticket_status"`   // 车票状态 0:有效；1:无效
	Departure      string    `gorm:"column:departure"`       // 出发站点
	Arrival        string    `gorm:"column:arrival"`         // 到达站点
	DepartureTime  time.Time `gorm:"column:departure_time"`  // 出发时间 yyyy-MM-dd HH:mm:ss
	ArrivalTime    time.Time `gorm:"column:arrival_time"`    // 到达时间 yyyy-MM-dd HH:mm:ss
}

// TableName 设置表名
func (TicketDO) TableName() string {
	return "t_ticket"
}

const (
	TicketValid   = 0 // 车票有效
	TicketInvalid = 1 // 车票无效
)

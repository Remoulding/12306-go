package model

// TicketDO 车票实体
type TicketDO struct {
	BaseDO                // 嵌入 BaseDO 模拟继承
	ID             int64  `gorm:"primaryKey;column:id"`   // id
	Username       string `gorm:"column:username"`        // 用户名
	TrainID        int64  `gorm:"column:train_id"`        // 列车id
	CarriageNumber string `gorm:"column:carriage_number"` // 车厢号
	SeatNumber     string `gorm:"column:seat_number"`     // 座位号
	PassengerID    string `gorm:"column:passenger_id"`    // 乘车人 ID
	TicketStatus   int    `gorm:"column:ticket_status"`   // 车票状态
}

// TableName 设置表名
func (TicketDO) TableName() string {
	return "t_ticket"
}

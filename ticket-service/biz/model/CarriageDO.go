package model

// CarriageDO 车厢实体
type CarriageDO struct {
	BaseDO              // 嵌入 BaseDO 模拟继承
	ID           int64  `gorm:"primaryKey;column:id"`   // id
	TrainID      int64  `gorm:"column:train_id"`        // 列车ID
	CarriageNum  string `gorm:"column:carriage_number"` // 车厢号
	CarriageType int    `gorm:"column:carriage_type"`   // 车厢类型
	SeatCount    int    `gorm:"column:seat_count"`      // 座位数
}

// TableName 设置表名
func (CarriageDO) TableName() string {
	return "t_carriage"
}

package model

// TrainDO 列车实体
type TrainDO struct {
	BaseDO               // 嵌入 BaseDO 模拟继承
	ID            int64  `gorm:"primaryKey;column:id"`  // id
	TrainNumber   string `gorm:"column:train_number"`   // 列车车次
	StartStation  string `gorm:"column:start_station"`  // 起始站
	EndStation    string `gorm:"column:end_station"`    // 终点站
	StartRegion   string `gorm:"column:start_region"`   // 起始城市
	EndRegion     string `gorm:"column:end_region"`     // 终点城市
	DepartureTime string `gorm:"column:departure_time"` // 出发时间 HH:MM:SS
	ArrivalTime   string `gorm:"column:arrival_time"`   // 到达时间 HH:MM:SS
	CrossDay      int    `gorm:"column:cross_day"`      // 跨越的天数
}

// TableName 设置表名
func (TrainDO) TableName() string {
	return "t_train"
}

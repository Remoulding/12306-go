package model

// TrainStationDO 列车车站实体
type TrainStationDO struct {
	BaseDO               // 嵌入 BaseDO 模拟继承
	ID            int64  `gorm:"primaryKey;column:id"`  // id
	TrainID       int64  `gorm:"column:train_id"`       // 车次id
	StationID     int64  `gorm:"column:station_id"`     // 车站id
	Sequence      string `gorm:"column:sequence"`       // 站点顺序
	Departure     string `gorm:"column:departure"`      // 出发站点
	Arrival       string `gorm:"column:arrival"`        // 到达站点
	StartRegion   string `gorm:"column:start_region"`   // 起始城市
	EndRegion     string `gorm:"column:end_region"`     // 终点城市
	ArrivalTime   string `gorm:"column:arrival_time"`   // 到站时间
	DepartureTime string `gorm:"column:departure_time"` // 出站时间
	StopoverTime  int    `gorm:"column:stopover_time"`  // 停留时间，单位分
	CrossDay      int    `gorm:"column:cross_day"`      // 跨越的天数
}

// TableName 设置表名
func (TrainStationDO) TableName() string {
	return "t_train_station"
}

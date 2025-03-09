package model

import "time"

// TrainDO 列车实体
type TrainDO struct {
	BaseDO                  // 嵌入 BaseDO 模拟继承
	ID            int64     `gorm:"primaryKey;column:id"`  // id
	TrainNumber   string    `gorm:"column:train_number"`   // 列车车次
	TrainType     int       `gorm:"column:train_type"`     // 列车类型 0：高铁 1：动车 2：普通车
	TrainTag      string    `gorm:"column:train_tag"`      // 列车标签 0：复兴号 1：智能动车组 2：静音车厢 3：支持选铺
	TrainBrand    string    `gorm:"column:train_brand"`    // 列车品牌类型
	StartStation  string    `gorm:"column:start_station"`  // 起始站
	EndStation    string    `gorm:"column:end_station"`    // 终点站
	StartRegion   string    `gorm:"column:start_region"`   // 起始城市
	EndRegion     string    `gorm:"column:end_region"`     // 终点城市
	SaleTime      time.Time `gorm:"column:sale_time"`      // 销售时间
	SaleStatus    int       `gorm:"column:sale_status"`    // 销售状态 0：可售 1：不可售 2：未知
	DepartureTime time.Time `gorm:"column:departure_time"` // 出发时间
	ArrivalTime   time.Time `gorm:"column:arrival_time"`   // 到达时间
}

// TableName 设置表名
func (TrainDO) TableName() string {
	return "t_train"
}

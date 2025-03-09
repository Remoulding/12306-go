package model

// StationDO 车站实体
type StationDO struct {
	BaseDO            // 嵌入 BaseDO 模拟继承
	ID         int64  `gorm:"primaryKey;column:id"` // id
	Code       string `gorm:"column:code"`          // 车站编码
	Name       string `gorm:"column:name"`          // 车站名称
	Spell      string `gorm:"column:spell"`         // 拼音
	Region     string `gorm:"column:region"`        // 地区编号
	RegionName string `gorm:"column:region_name"`   // 地区名称
}

// TableName 设置表名
func (StationDO) TableName() string {
	return "t_station"
}

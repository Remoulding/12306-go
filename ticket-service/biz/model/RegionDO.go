package model

// RegionDO 地区实体
type RegionDO struct {
	BaseDO             // 嵌入 BaseDO 模拟继承
	ID          int64  `gorm:"primaryKey;column:id"` // id
	Name        string `gorm:"column:name"`          // 地区名称
	FullName    string `gorm:"column:full_name"`     // 地区全名
	Code        string `gorm:"column:code"`          // 地区编码
	Initial     string `gorm:"column:initial"`       // 地区首字母
	Spell       string `gorm:"column:spell"`         // 拼音
	PopularFlag int    `gorm:"column:popular_flag"`  // 热门标识
}

// TableName 设置表名
func (RegionDO) TableName() string {
	return "t_region"
}

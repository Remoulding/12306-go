package model

import (
	"time"
)

// BaseDO 是数据持久层基础属性
type BaseDO struct {
	// 创建时间，自动填充
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime;"`
	// 修改时间，自动填充
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime;"`
	// 删除标志，插入时自动填充
	DelFlag int `gorm:"column:del_flag;default:0;"`
}

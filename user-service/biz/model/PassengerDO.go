package model

import "time"

// PassengerDO 乘车人实体
type PassengerDO struct {
	BaseDO                  // 嵌入 BaseDO 模拟继承
	ID           uint64     `gorm:"primaryKey;column:id"` // id
	Username     string     `gorm:"column:username"`      // 用户名
	RealName     string     `gorm:"column:real_name"`     // 真实姓名
	IDType       int        `gorm:"column:id_type"`       // 证件类型
	IDCard       string     `gorm:"column:id_card"`       // 证件号码
	DiscountType int        `gorm:"column:discount_type"` // 优惠类型
	Phone        string     `gorm:"column:phone"`         // 手机号
	CreateDate   *time.Time `gorm:"column:create_date"`   // 添加日期
	VerifyStatus int        `gorm:"column:verify_status"` // 审核状态
}

// TableName 设置表名
func (PassengerDO) TableName() string {
	return "t_passenger"
}

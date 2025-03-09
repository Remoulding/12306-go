package model

// UserDO 用户信息实体
type UserDO struct {
	BaseDO              // 嵌入 BaseDO 模拟继承
	ID           uint64 `gorm:"primaryKey;column:id"` // id
	Username     string `gorm:"column:username"`      // 用户名
	Password     string `gorm:"column:password"`      // 密码
	RealName     string `gorm:"column:real_name"`     // 真实姓名
	Region       string `gorm:"column:region"`        // 国家/地区
	IDType       int    `gorm:"column:id_type"`       // 证件类型
	IDCard       string `gorm:"column:id_card"`       // 证件号
	Phone        string `gorm:"column:phone"`         // 手机号
	Telephone    string `gorm:"column:telephone"`     // 固定电话
	Mail         string `gorm:"column:mail"`          // 邮箱
	UserType     int    `gorm:"column:user_type"`     // 旅客类型
	VerifyStatus int    `gorm:"column:verify_status"` // 审核状态
	PostCode     string `gorm:"column:post_code"`     // 邮编
	Address      string `gorm:"column:address"`       // 地址
	DeletionTime int64  `gorm:"column:deletion_time"` // 注销时间戳
}

// TableName 设置表名
func (UserDO) TableName() string {
	return "t_user"
}

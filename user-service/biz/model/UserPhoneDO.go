package model

// UserPhoneDO 用户手机号实体，继承 BaseDO
type UserPhoneDO struct {
	BaseDO
	ID           int64  `json:"id"`            // id
	Username     string `json:"username"`      // 用户名
	Phone        string `json:"phone"`         // 手机号
	DeletionTime int64  `json:"deletion_time"` // 注销时间戳
}

// TableName 实现 GORM 的 TableName 方法来指定表名
func (UserPhoneDO) TableName() string {
	return "t_user_phone" // 指定表名
}

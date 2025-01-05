package model

// UserMailDO 用户邮箱表实体，继承 BaseDO
type UserMailDO struct {
	BaseDO
	ID           int64  `json:"id"`            // id
	Username     string `json:"username"`      // 用户名
	Mail         string `json:"mail"`          // 手机号
	DeletionTime int64  `json:"deletion_time"` // 注销时间戳
}

// TableName 实现 GORM 的 TableName 方法来指定表名
func (UserMailDO) TableName() string {
	return "t_user_mail" // 指定表名
}

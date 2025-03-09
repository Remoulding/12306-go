package model

type UserReuseDO struct {
	BaseDO
	Username string `json:"username"` // 用户名
}

// TableName 实现 GORM 的 TableName 方法来指定表名
func (UserReuseDO) TableName() string {
	return "t_user_reuse" // 指定表名
}

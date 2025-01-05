package model

// UserDeletionDO 用户注销表实体，继承 BaseDO
type UserDeletionDO struct {
	BaseDO
	ID     int64  `json:"id"`      // id
	IdType int    `json:"id_type"` // 证件类型
	IdCard string `json:"id_card"` // 证件号
}

// TableName 实现 GORM 的 TableName 方法来指定表名
func (UserDeletionDO) TableName() string {
	return "t_user_deletion" // 指定表名
}

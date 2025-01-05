package req

// UserUpdateReqDTO 用户修改请求参数
type UserUpdateReqDTO struct {
	// 用户ID
	ID string `json:"id"`

	// 用户名
	Username string `json:"username"`

	// 邮箱
	Mail string `json:"mail"`

	// 旅客类型
	UserType int `json:"userType"`

	// 邮编
	PostCode string `json:"postCode"`

	// 地址
	Address string `json:"address"`
}

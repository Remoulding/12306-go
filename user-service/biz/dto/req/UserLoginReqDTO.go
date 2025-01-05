package req

// UserLoginReqDTO 用户登录请求参数
type UserLoginReqDTO struct {
	// 用户名、邮箱或手机号
	UsernameOrMailOrPhone string `json:"usernameOrMailOrPhone"`

	// 密码
	Password string `json:"password"`
}

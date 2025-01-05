package resp

type UserRegisterRespDTO struct {
	// 用户名
	Username string `json:"username"`
	// 真实姓名
	RealName string `json:"realName"`
	// 手机号
	Phone string `json:"phone"`
}

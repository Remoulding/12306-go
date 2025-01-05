package req

// UserRegisterReqDTO 表示用户注册的请求参数
type UserRegisterReqDTO struct {
	// 用户名
	Username string `json:"username"`

	// 密码
	Password string `json:"password"`

	// 真实姓名
	RealName string `json:"realName"`

	// 证件类型
	IdType int `json:"idType"`

	// 证件号
	IdCard string `json:"idCard"`

	// 手机号
	Phone string `json:"phone"`

	// 邮箱
	Mail string `json:"mail"`

	// 旅客类型
	UserType int `json:"userType"`

	// 审核状态
	VerifyState int `json:"verifyState"`

	// 邮编
	PostCode string `json:"postCode"`

	// 地址
	Address string `json:"address"`

	// 国家/地区
	Region string `json:"region"`

	// 固定电话
	Telephone string `json:"telephone"`
}

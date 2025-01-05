package resp

// UserLoginRespDTO represents the response DTO for user login.
type UserLoginRespDTO struct {
	// 用户 ID
	UserId string `json:"userId"`

	// 用户名
	Username string `json:"username"`

	// 真实姓名
	RealName string `json:"realName"`

	// Token
	AccessToken string `json:"accessToken"`
}

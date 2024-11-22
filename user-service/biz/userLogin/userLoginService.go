package userLogin

type UserLoginReqDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRespDTO struct {
	UserID      string `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type Result[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
}

func Success[T any](data *T) *Result[T] {
	return &Result[T]{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
}

type UserLoginService struct{}

func (s *UserLoginService) Login(req UserLoginReqDTO) UserLoginRespDTO {
	// 模拟返回数据
	return UserLoginRespDTO{
		UserID:      "12345",
		AccessToken: "abc123token",
	}
}

func (s *UserLoginService) CheckLogin(accessToken string) UserLoginRespDTO {
	// 模拟检查逻辑
	return UserLoginRespDTO{
		UserID:      "12345",
		AccessToken: accessToken,
	}
}

func (s *UserLoginService) Logout(accessToken string) {
	// 模拟注销逻辑
}

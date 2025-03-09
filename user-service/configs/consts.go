package configs

import "time"

// 常量定义
const (
	CacheTTL = 60 * time.Minute
)

const (
	// UserIdKey 用户 ID Key
	UserIdKey = "UserId"

	// UserNameKey 用户名 Key
	UserNameKey = "Username"

	// RealNameKey 用户真实名称 Key
	RealNameKey = "RealName"

	// UserTokenKey 用户 Token Key
	UserTokenKey = "Token"
)

const (
	// JwtSecretKey JWT 密钥
	JwtSecretKey = "12306-jwt-secret-key"
)

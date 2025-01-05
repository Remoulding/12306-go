package resp

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"strings"
)

// UserQueryRespDTO 用户查询返回参数
type UserQueryRespDTO struct {
	Username     string `json:"username"`
	RealName     string `json:"real_name"`
	Region       string `json:"region"`
	IdType       int    `json:"id_type"`
	IdCard       string `json:"id_card"`
	Phone        string `json:"phone"`
	Telephone    string `json:"telephone"`
	Mail         string `json:"mail"`
	UserType     int    `json:"user_type"`
	VerifyStatus int    `json:"verify_status"`
	PostCode     string `json:"post_code"`
	Address      string `json:"address"`
}

// MarshalJSON Custom serializer for ID Card to mask sensitive information
func (u *UserQueryRespDTO) MarshalJSON() ([]byte, error) {
	type Alias UserQueryRespDTO
	return json.Marshal(&struct {
		IdCard string `json:"id_card"`
		Phone  string `json:"phone"`
		*Alias
	}{
		IdCard: maskIdCard(u.IdCard),
		Phone:  maskPhone(u.Phone),
		Alias:  (*Alias)(u),
	})
}

// maskIdCard 用于脱敏证件号
func maskIdCard(idCard string) string {
	// 假设我们只保留证件号的前六位和后四位，其他部分使用 * 代替
	if len(idCard) < 10 {
		return idCard
	}
	return idCard[:6] + strings.Repeat("*", len(idCard)-10) + idCard[len(idCard)-4:]
}

// maskPhone 用于脱敏手机号
func maskPhone(phone string) string {
	// 假设我们只保留手机号的前三位和后四位，其他部分使用 * 代替
	if len(phone) < 10 {
		return phone
	}
	return phone[:3] + strings.Repeat("*", len(phone)-7) + phone[len(phone)-4:]
}

// 用于验证字段的合法性
var validate = validator.New()

// Validate 使用 go-playground/validator 库对结构体进行验证
func (u *UserQueryRespDTO) Validate() error {
	return validate.Struct(u)
}

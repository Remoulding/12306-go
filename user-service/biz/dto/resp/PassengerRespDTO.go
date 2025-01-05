package resp

import (
	"encoding/json"
	"time"
)

// PassengerRespDTO 乘车人返回参数
type PassengerRespDTO struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	RealName     string    `json:"realName"`
	IDType       int       `json:"idType"`
	IDCard       string    `json:"idCard"`
	ActualIDCard string    `json:"actualIdCard"`
	DiscountType int       `json:"discountType"`
	Phone        string    `json:"phone"`
	ActualPhone  string    `json:"actualPhone"`
	CreateDate   time.Time `json:"createDate"`
	VerifyStatus int       `json:"verifyStatus"`
}

// PhoneDesensitizationSerializer 处理手机号脱敏的结构体
type PhoneDesensitizationSerializer struct {
	Phone string
}

// IdCardDesensitizationSerializer 处理身份证号脱敏的结构体
type IdCardDesensitizationSerializer struct {
	IDCard string
}

// serialize 处理手机号脱敏
func (p PhoneDesensitizationSerializer) serialize() string {
	// 只返回手机号的前三位和后四位，其他部分用星号替代
	return p.Phone[:3] + "****" + p.Phone[len(p.Phone)-4:]
}

// serialize 处理身份证号脱敏
func (i IdCardDesensitizationSerializer) serialize() string {
	// 只返回身份证号的前六位和后四位，其他部分用星号替代
	return i.IDCard[:6] + "****" + i.IDCard[len(i.IDCard)-4:]
}

// MarshalJSON 重写序列化方法处理脱敏字段
func (p *PassengerRespDTO) MarshalJSON() ([]byte, error) {
	type Alias PassengerRespDTO
	return json.Marshal(&struct {
		Phone      string `json:"phone"`
		IDCard     string `json:"idCard"`
		CreateDate string `json:"createDate"`
		*Alias
	}{
		Phone:      PhoneDesensitizationSerializer{Phone: p.Phone}.serialize(),
		IDCard:     IdCardDesensitizationSerializer{IDCard: p.IDCard}.serialize(),
		CreateDate: p.CreateDate.Format("2006-01-02"),
		Alias:      (*Alias)(p),
	})
}

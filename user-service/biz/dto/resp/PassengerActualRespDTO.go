package resp

import (
	"encoding/json"
	"time"
)

// PassengerActualRespDTO 乘车人真实返回参数，不包含脱敏信息
type PassengerActualRespDTO struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	RealName     string    `json:"realName"`
	IDType       int       `json:"idType"`
	IDCard       string    `json:"idCard"`
	DiscountType int       `json:"discountType"`
	Phone        string    `json:"phone"`
	CreateDate   time.Time `json:"createDate"`
	VerifyStatus int       `json:"verifyStatus"`
}

// MarshalJSON 自定义 JSON 序列化方法，用于日期格式化
func (p *PassengerActualRespDTO) MarshalJSON() ([]byte, error) {
	type Alias PassengerActualRespDTO
	return json.Marshal(&struct {
		CreateDate string `json:"createDate"`
		*Alias
	}{
		CreateDate: p.CreateDate.Format("2006-01-02"),
		Alias:      (*Alias)(p),
	})
}

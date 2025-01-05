package req

// PassengerReqDTO 乘车人添加&修改请求参数
type PassengerReqDTO struct {
	ID           string `json:"id"`
	RealName     string `json:"realName"`
	IDType       int    `json:"idType"`
	IDCard       string `json:"idCard"`
	DiscountType int    `json:"discountType"`
	Phone        string `json:"phone"`
}

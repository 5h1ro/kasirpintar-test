package request

type StatusRequest struct {
	TimeStamp     string `json:"timeStamp"`
	TXid          string `json:"tXid" binding:"required" example:"IONPAYTEST01202307131229186508"`
	IMid          string `json:"iMid"`
	ReferenceNo   string `json:"referenceNo" binding:"required" example:"TESTING20230713052918"`
	Amt           string `json:"amt" binding:"required" example:"1000"`
	MerchantToken string `json:"merchantToken"`
}

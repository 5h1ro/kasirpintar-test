package request

type PaymentRequest struct {
	ReferenceNo    string `json:"referenceNo" binding:"required" example:"TESTING20230713052918"`
	Amt            string `json:"amt" binding:"required" example:"1000"`
	TimeStamp      string `json:"timeStamp"`
	TXid           string `json:"tXid" binding:"required" example:"IONPAYTEST01202307131229186508"`
	CardNo         string `json:"cardNo" binding:"required" example:"4111111111111111"`
	CardExpYymm    string `json:"cardExpYymm" binding:"required" example:"2512"`
	CardCvv        string `json:"cardCvv" binding:"required" example:"123"`
	CardHolderNm   string `json:"cardHolderNm" binding:"required" example:"JOHN DOE"`
	RecurringToken string `json:"recurringToken"`
	PreauthToken   string `json:"preauthToken"`
	MerchantToken  string `json:"merchantToken"`
	CallBackUrl    string `json:"callBackUrl"`
	ClickPayNo     string `json:"clickPayNo"`
	ClickPayToken  string `json:"clickPayToken"`
	DataField3     string `json:"dataField3"`
}

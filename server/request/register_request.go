package request

type RegisterRequest struct {
	IMid            string `json:"iMid"`
	MerchantToken   string `json:"merchantToken"`
	PayMethod       string `json:"payMethod" binding:"required" example:"01"`
	Currency        string `json:"currency" binding:"required" example:"IDR"`
	Amt             string `json:"amt" binding:"required" example:"1000"`
	ReferenceNo     string `json:"referenceNo"`
	GoodsNm         string `json:"goodsNm" binding:"required" example:"Merchant Goods 1"`
	BillingNm       string `json:"billingNm" binding:"required" example:"John Doe"`
	BillingPhone    string `json:"billingPhone" binding:"required" example:"2112345678"`
	BillingEmail    string `json:"billingEmail" binding:"required" example:"buyer@merchant.com"`
	BillingAddr     string `json:"billingAddr" binding:"required" example:"Jln Merdeka 123"`
	BillingCity     string `json:"billingCity" binding:"required" example:"Jakara Selatan"`
	BillingState    string `json:"billingState" binding:"required" example:"DKI Jakarta"`
	BillingPostCd   string `json:"billingPostCd" binding:"required" example:"14350"`
	BillingCountry  string `json:"billingCountry" binding:"required" example:"Indonesia"`
	DeliveryNm      string `json:"deliveryNm" example:"John Doe"`
	DeliveryPhone   string `json:"deliveryPhone" example:"2112345678"`
	DeliveryAddr    string `json:"deliveryAddr" example:"Jln Merdeka 123"`
	DeliveryCity    string `json:"deliveryCity" example:"Jakara Selatan"`
	DeliveryState   string `json:"deliveryState" example:"DKI Jakarta"`
	DeliveryPostCd  string `json:"deliveryPostCd" example:"14350"`
	DeliveryCountry string `json:"deliveryCountry" example:"Indonesia"`
	DbProcessUrl    string `json:"dbProcessUrl"`
	Vat             string `json:"vat"`
	Fee             string `json:"fee"`
	NotaxAmt        string `json:"notaxAmt"`
	Description     string `json:"description"`
	ReqDt           string `json:"reqDt"`
	ReqTm           string `json:"reqTm"`
	ReqServerIP     string `json:"reqServerIP"`
	ReqClientVer    string `json:"reqClientVer"`
	ReqDomain       string `json:"reqDomain"`
	UserIP          string `json:"userIP" binding:"required" example:"127.0.0.1"`
	UserSessionID   string `json:"userSessionID"`
	UserAgent       string `json:"userAgent"`
	UserLanguage    string `json:"userLanguage"`
	CartData        string `json:"cartData" binding:"required" example:"{}"`
	InstmntType     string `json:"instmntType" binding:"required" example:"1"`
	InstmntMon      string `json:"instmntMon" binding:"required" example:"1"`
	RecurrOpt       string `json:"recurrOpt" binding:"required" example:"2"`
	VacctValidDt    string `json:"vacctValidDt"`
	VacctValidTm    string `json:"vacctValidTm"`
	MerFixAcctId    string `json:"merFixAcctId" example:"4"`
	MitraCd         string `json:"mitraCd"`
	BankCd          string `json:"bankCd"`
	TimeStamp       string `json:"timeStamp"`
}

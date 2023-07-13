package response

type PaymentResponse struct {
	ResultCd       string `json:"resultCd"`
	ResultMsg      string `json:"resultMsg"`
	TXid           string `json:"tXid"`
	ReferenceNo    string `json:"referenceNo"`
	PayMethod      string `json:"payMethod"`
	Amt            string `json:"amt"`
	TransDt        string `json:"transDt"`
	TransTm        string `json:"transTm"`
	Description    string `json:"description"`
	AuthNo         string `json:"authNo"`
	IssuBankCd     string `json:"issuBankCd"`
	AcquBankCd     string `json:"acquBankCd"`
	CardNo         string `json:"cardNo"`
	ReceiptCode    string `json:"receiptCode"`
	MitraCd        string `json:"mitraCd"`
	RecurringToken string `json:"recurringToken"`
	PreauthToken   string `json:"preauthToken"`
	Currency       string `json:"currency"`
	GoodsNm        string `json:"goodsNm"`
	BillingNm      string `json:"billingNm"`
	CcTransType    string `json:"ccTransType"`
	MRefNo         string `json:"mRefNo"`
	InstmntType    string `json:"instmntType"`
	InstmntMon     string `json:"instmntMon"`
	CardExpYymm    string `json:"cardExpYymm"`
	IssuBankNm     string `json:"issuBankNm"`
	AcquBankNm     string `json:"acquBankNm"`
	TimeStamp      string `json:"timeStamp"`
	MerchantToken  string `json:"merchantToken"`
	CardBrand      string `json:"cardBrand"`
}

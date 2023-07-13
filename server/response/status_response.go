package response

type StatusResponse struct {
	ResultCd    string `json:"resultCd"`
	ResultMsg   string `json:"resultMsg"`
	TXid        string `json:"tXid"`
	IMid        string `json:"iMid"`
	ReferenceNo string `json:"referenceNo"`
	PayMethod   string `json:"payMethod"`
	Amt         string `json:"amt"`
	CancelAmt   string `json:"cancelAmt"`
	ReqDt       string `json:"reqDt"`
	ReqTm       string `json:"reqTm"`
	TransDt     string `json:"transDt"`
	TransTm     string `json:"transTm"`
	DepositDt   string `json:"depositDt"`
	DepositTm   string `json:"depositTm"`
	Currency    string `json:"currency"`
	GoodsNm     string `json:"goodsNm"`
	BillingNm   string `json:"billingNm"`
	Status      string `json:"status"`
}

package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"kaspin/server/request"
	"kaspin/server/response"
	"kaspin/server/utils"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type TransactionService struct{}

func (service TransactionService) Register(r request.RegisterRequest, reqLog, resLog zerolog.Logger) (*response.RegisterResponse, error) {
	timestamp := time.Now().Format("20060102150405")
	merchantID := os.Getenv("MERCHANTID")
	merchantKey := os.Getenv("MERCHANTKEY")
	merchantDbProcess := os.Getenv("MERCHANTDBPROCESS")
	referenceNo := "TESTING" + timestamp
	endpoint := fmt.Sprintf("%s/nicepay/direct/v2/registration", os.Getenv("NICEPAYHOST"))
	// generate token
	token, _ := utils.MerchatTokenGenerator(timestamp, merchantID, referenceNo, r.Amt, merchantKey)
	byteHashed := sha256.Sum256([]byte(token))
	hashedToken := hex.EncodeToString(byteHashed[:])
	// fill request
	r.DbProcessUrl = merchantDbProcess
	r.ReferenceNo = referenceNo
	r.MerchantToken = hashedToken
	r.IMid = merchantID
	r.TimeStamp = timestamp

	// convert data to bytes
	payloadBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqLog.Info().RawJSON("payload", payloadBytes).Msg("Register request data (system)")
	// initiate request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	// set header
	req.Header.Set("Content-Type", "application/json")
	// send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// handle response data
	resBody := new(bytes.Buffer)
	_, err = resBody.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}
	responseData := &response.RegisterResponse{}
	err = json.Unmarshal(resBody.Bytes(), &responseData)
	if err != nil {
		return nil, err
	}
	resLog.Info().RawJSON("payload", resBody.Bytes()).Msg("Register response data (system)")

	return responseData, nil
}

func (service TransactionService) Status(r request.StatusRequest, reqLog, resLog zerolog.Logger) (*response.StatusResponse, error) {
	timestamp := time.Now().Format("20060102150405")
	merchantID := os.Getenv("MERCHANTID")
	merchantKey := os.Getenv("MERCHANTKEY")
	endpoint := fmt.Sprintf("%s/nicepay/direct/v2/inquiry", os.Getenv("NICEPAYHOST"))
	// generate token
	token, _ := utils.MerchatTokenGenerator(timestamp, merchantID, r.ReferenceNo, r.Amt, merchantKey)
	byteHashed := sha256.Sum256([]byte(token))
	hashedToken := hex.EncodeToString(byteHashed[:])
	// fill request
	r.MerchantToken = hashedToken
	r.IMid = merchantID
	r.TimeStamp = timestamp
	// convert data to bytes
	payloadBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqLog.Info().RawJSON("payload", payloadBytes).Msg("Check status request data (system)")
	// initiate request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	// set header
	req.Header.Set("Content-Type", "application/json")
	// send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// handle response data
	resBody := new(bytes.Buffer)
	_, err = resBody.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}
	responseData := &response.StatusResponse{}
	err = json.Unmarshal(resBody.Bytes(), &responseData)
	if err != nil {
		return nil, err
	}
	resLog.Info().RawJSON("payload", resBody.Bytes()).Msg("Check status response data (system)")
	return responseData, nil
}

func (service TransactionService) Payment(r request.PaymentRequest, reqLog, resLog zerolog.Logger) (*response.PaymentResponse, error) {
	timestamp := time.Now().Format("20060102150405")
	merchantID := os.Getenv("MERCHANTID")
	merchantKey := os.Getenv("MERCHANTKEY")
	merchantCallback := os.Getenv("MERCHANTCALLBACK")
	endpoint := fmt.Sprintf("%s/nicepay/direct/v2/payment", os.Getenv("NICEPAYHOST"))
	// generate token
	token, _ := utils.MerchatTokenGenerator(timestamp, merchantID, r.ReferenceNo, r.Amt, merchantKey)
	byteHashed := sha256.Sum256([]byte(token))
	hashedToken := hex.EncodeToString(byteHashed[:])
	// fill request
	r.MerchantToken = hashedToken
	r.TimeStamp = timestamp
	r.CallBackUrl = merchantCallback
	// convert data to bytes
	payloadBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	reqLog.Info().RawJSON("payload", payloadBytes).Msg("Payment request data (system)")
	// preparing data
	data := url.Values{
		"timeStamp":      {r.TimeStamp},
		"tXid":           {r.TXid},
		"cardNo":         {r.CardNo},
		"cardExpYymm":    {r.CardExpYymm},
		"cardCvv":        {r.CardCvv},
		"cardHolderNm":   {r.CardHolderNm},
		"recurringToken": {r.RecurringToken},
		"preauthToken":   {r.PreauthToken},
		"merchantToken":  {r.MerchantToken},
		"callBackUrl":    {r.CallBackUrl},
		"clickPayNo":     {r.ClickPayNo},
		"dataField3":     {r.DataField3},
		"clickPayToken":  {r.ClickPayToken},
	}.Encode()
	// initiate request
	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	// set header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
	// send request
	client := &http.Client{}
	res, err := client.Do(req)
	// handle response
	respDump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}
	// extract data
	params := []string{
		"resultCd",
		"resultMsg",
		"tXid",
		"referenceNo",
		"payMethod",
		"amt",
		"transDt",
		"transTm",
		"description",
		"authNo",
		"issuBankCd",
		"acquBankCd",
		"cardNo",
		"receiptCode",
		"mitraCd",
		"recurringToken",
		"preauthToken",
		"currency",
		"goodsNm",
		"billingNm",
		"ccTransType",
		"mRefNo",
		"instmntType",
		"instmntMon",
		"cardExpYymm",
		"issuBankNm",
		"acquBankNm",
		"timeStamp",
		"merchantToken",
		"cardBrand",
	}
	var result = response.PaymentResponse{}
	for _, v := range params {
		start := strings.Index(string(respDump), v)
		value := string(respDump)[start:]
		startIndex := strings.Index(value, `value="`)
		value = value[startIndex+7:]
		endIndex := strings.Index(value, `"`)
		value = value[:endIndex]
		switch v {
		case "resultCd":
			result.ResultCd = value
		case "resultMsg":
			result.ResultMsg = value
		case "tXid":
			result.TXid = value
		case "referenceNo":
			result.ReferenceNo = value
		case "payMethod":
			result.PayMethod = value
		case "amt":
			result.Amt = value
		case "transDt":
			result.TransDt = value
		case "transTm":
			result.TransTm = value
		case "description":
			result.Description = value
		case "authNo":
			result.AuthNo = value
		case "issuBankCd":
			result.IssuBankCd = value
		case "acquBankCd":
			result.AcquBankCd = value
		case "cardNo":
			result.CardNo = value
		case "receiptCode":
			result.ReceiptCode = value
		case "mitraCd":
			result.MitraCd = value
		case "recurringToken":
			result.RecurringToken = value
		case "preauthToken":
			result.PreauthToken = value
		case "currency":
			result.Currency = value
		case "goodsNm":
			result.GoodsNm = value
		case "billingNm":
			result.BillingNm = value
		case "ccTransType":
			result.CcTransType = value
		case "mRefNo":
			result.MRefNo = value
		case "instmntType":
			result.InstmntType = value
		case "instmntMon":
			result.InstmntMon = value
		case "cardExpYymm":
			result.CardExpYymm = value
		case "issuBankNm":
			result.IssuBankNm = value
		case "acquBankNm":
			result.AcquBankNm = value
		case "timeStamp":
			result.TimeStamp = value
		case "merchantToken":
			result.MerchantToken = value
		case "cardBrand":
			result.CardBrand = value
		}
	}
	// logging
	resBytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	resLog.Info().RawJSON("payload", resBytes).Msg("Payment response data (system)")
	return &result, nil
}

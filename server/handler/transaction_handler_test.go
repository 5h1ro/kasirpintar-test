package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kaspin/server/response"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

var tXid, referenceNo string
var reqLog, resLog, errLog zerolog.Logger

type TestTransactionSuite struct {
	suite.Suite
}

func (suite *TestTransactionSuite) SetupSuite() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	println(">>> START Setup suite")
}

func (suite *TestTransactionSuite) TearDownTest() {
	println(">>> END   Test complete")
}

// Unit test register
func (suite *TestTransactionSuite) Test1NegativeRegister() {
	println(">>> START Negative test register (validation check) [-]")

	gin.SetMode(gin.TestMode)

	server := gin.New()
	server.POST(
		"/register",
		NewTransactionHandler(reqLog, resLog, errLog).Register(),
	)

	recorder := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer([]byte(`{"billingAddr": "Jln Merdeka 123","billingCity": "Jakara Selatan","billingCountry": "Indonesia","billingEmail": "buyer@merchant.com","billingNm": "John Doe","billingPhone": "2112345678","billingPostCd": "14350","billingState": "DKI Jakarta","cartData": "{}","currency": "IDR","goodsNm": "Merchant Goods 1","instmntMon": "1","instmntType": "1","payMethod": "01","recurrOpt": "2","userIP": "127.0.0.1"}`)),
	)
	req.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(recorder, req)

	suite.Equal(http.StatusUnprocessableEntity, recorder.Code)
}
func (suite *TestTransactionSuite) Test2NegativeRegister() {
	println(">>> START Negative test register (Error from Nicepay) [-]")

	gin.SetMode(gin.TestMode)

	server := gin.New()
	server.POST(
		"/register",
		NewTransactionHandler(reqLog, resLog, errLog).Register(),
	)

	recorder := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer([]byte(`{"amt": "1000","billingAddr": "Jln Merdeka 123","billingCity": "Jakara Selatan","billingCountry": "Indonesia","billingEmail": "buyer@merchant.com","billingNm": "John Doe","billingPhone": "2112345678","billingPostCd": "14350","billingState": "DKI Jakarta","cartData": "{}","currency": "IDR","goodsNm": "Merchant Goods 1","instmntMon": "1","instmntType": "1","payMethod": "02","recurrOpt": "2","userIP": "127.0.0.1"}`)),
	)
	req.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(recorder, req)

	suite.Equal(http.StatusBadRequest, recorder.Code)
}
func (suite *TestTransactionSuite) Test3PositiveRegister() {
	println(">>> START Positive test register (Successfull register) [+]")

	gin.SetMode(gin.TestMode)

	server := gin.New()
	server.POST(
		"/register",
		NewTransactionHandler(reqLog, resLog, errLog).Register(),
	)

	recorder := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer([]byte(`{"amt": "1000","billingAddr": "Jln Merdeka 123","billingCity": "Jakara Selatan","billingCountry": "Indonesia","billingEmail": "buyer@merchant.com","billingNm": "John Doe","billingPhone": "2112345678","billingPostCd": "14350","billingState": "DKI Jakarta","cartData": "{}","currency": "IDR","goodsNm": "Merchant Goods 1","instmntMon": "1","instmntType": "1","payMethod": "01","recurrOpt": "2","userIP": "127.0.0.1"}`)),
	)
	req.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(recorder, req)
	resBody := new(bytes.Buffer)
	_, err := resBody.ReadFrom(recorder.Body)
	if err != nil {
		println(err.Error())
	}
	responseData := &response.RegisterResponse{}
	err = json.Unmarshal(resBody.Bytes(), &responseData)
	if err != nil {
		println(err.Error())
	}
	tXid = responseData.TXid
	referenceNo = responseData.ReferenceNo
	suite.Equal(http.StatusOK, recorder.Code)
}

// Unit test check status before payment
func (suite *TestTransactionSuite) Test4NegativeStatus() {
	println(">>> START Negative test check status (validation check) [-]")

	gin.SetMode(gin.TestMode)

	server := gin.New()
	server.POST(
		"/status",
		NewTransactionHandler(reqLog, resLog, errLog).Status(),
	)

	recorder := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodPost,
		"/status",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"tXid": %s,"referenceNo": %s}`, tXid, referenceNo))),
	)
	req.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(recorder, req)

	suite.Equal(http.StatusUnprocessableEntity, recorder.Code)
}
func (suite *TestTransactionSuite) Test5NegativeStatus() {
	println(">>> START Negative test check status (Error from Nicepay) [-]")

	gin.SetMode(gin.TestMode)

	server := gin.New()
	server.POST(
		"/status",
		NewTransactionHandler(reqLog, resLog, errLog).Status(),
	)

	recorder := httptest.NewRecorder()
	data := fmt.Sprintf(`{"amt": "5000","tXid": "%s","referenceNo": "%s"}`, tXid, referenceNo)
	req := httptest.NewRequest(
		http.MethodPost,
		"/status",
		bytes.NewBuffer([]byte(data)),
	)
	req.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(recorder, req)

	suite.Equal(http.StatusBadRequest, recorder.Code)
}
func (suite *TestTransactionSuite) Test6PositiveStatus() {
	println(">>> START Positive test check status (Successfull check status) [+]")

	gin.SetMode(gin.TestMode)

	server := gin.New()
	server.POST(
		"/status",
		NewTransactionHandler(reqLog, resLog, errLog).Status(),
	)

	recorder := httptest.NewRecorder()
	data := fmt.Sprintf(`{"amt": "1000","tXid": "%s","referenceNo": "%s"}`, tXid, referenceNo)
	req := httptest.NewRequest(
		http.MethodPost,
		"/status",
		bytes.NewBuffer([]byte(data)),
	)
	req.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(recorder, req)
	resBody := new(bytes.Buffer)
	_, err := resBody.ReadFrom(recorder.Body)
	if err != nil {
		println(err.Error())
	}
	responseData := &response.StatusResponse{}
	err = json.Unmarshal(resBody.Bytes(), &responseData)
	if err != nil {
		println(err.Error())
	}
	suite.Equal(http.StatusOK, recorder.Code)
}

// Unit test Payment
func (suite *TestTransactionSuite) Test7NegativePayment() {
	println(">>> START Negative test payment (validation payment) [-]")

	gin.SetMode(gin.TestMode)

	server := gin.New()
	server.POST(
		"/payment",
		NewTransactionHandler(reqLog, resLog, errLog).Payment(),
	)

	recorder := httptest.NewRecorder()

	req := httptest.NewRequest(
		http.MethodPost,
		"/payment",
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"cardCvv": "123","cardExpYymm": "2512","cardHolderNm": "JOHN DOE","cardNo": "4111111111111111","tXid": %s,"referenceNo": %s}`, tXid, referenceNo))),
	)
	req.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(recorder, req)

	suite.Equal(http.StatusUnprocessableEntity, recorder.Code)
}
func (suite *TestTransactionSuite) Test8NegativePayment() {
	println(">>> START Negative test payment (Error from Nicepay) [-]")

	gin.SetMode(gin.TestMode)

	server := gin.New()
	server.POST(
		"/payment",
		NewTransactionHandler(reqLog, resLog, errLog).Payment(),
	)

	recorder := httptest.NewRecorder()
	data := fmt.Sprintf(`{"amt": "1000","cardCvv": "123","cardExpYymm": "2012","cardHolderNm": "JOHN DOE","cardNo": "4111111111111111","tXid": "%s","referenceNo": "%s"}`, tXid, referenceNo)
	req := httptest.NewRequest(
		http.MethodPost,
		"/payment",
		bytes.NewBuffer([]byte(data)),
	)
	req.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(recorder, req)

	suite.Equal(http.StatusBadRequest, recorder.Code)
}
func (suite *TestTransactionSuite) Test9PositivePayment() {
	println(">>> START Positive test payment (Successfull payment) [+]")

	gin.SetMode(gin.TestMode)

	server := gin.New()
	server.POST(
		"/payment",
		NewTransactionHandler(reqLog, resLog, errLog).Payment(),
	)

	recorder := httptest.NewRecorder()
	data := fmt.Sprintf(`{"amt": "1000","cardCvv": "123","cardExpYymm": "2512","cardHolderNm": "JOHN DOE","cardNo": "4111111111111111","tXid": "%s","referenceNo": "%s"}`, tXid, referenceNo)
	req := httptest.NewRequest(
		http.MethodPost,
		"/payment",
		bytes.NewBuffer([]byte(data)),
	)
	req.Header.Add("Content-Type", "application/json")

	server.ServeHTTP(recorder, req)
	resBody := new(bytes.Buffer)
	_, err := resBody.ReadFrom(recorder.Body)
	if err != nil {
		println(err.Error())
	}
	responseData := &response.StatusResponse{}
	err = json.Unmarshal(resBody.Bytes(), &responseData)
	if err != nil {
		println(err.Error())
	}
	suite.Equal(http.StatusOK, recorder.Code)
}

func TestRunSuite(t *testing.T) {
	reqFile, err := os.OpenFile(
		"../../log/request.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	defer reqFile.Close()
	if err != nil {
		panic(err)
	}
	resFile, err := os.OpenFile(
		"../../log/response.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	defer resFile.Close()
	if err != nil {
		panic(err)
	}
	errFile, err := os.OpenFile(
		"../../log/error.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	defer errFile.Close()
	if err != nil {
		panic(err)
	}
	reqLog = zerolog.New(reqFile).With().Timestamp().Logger()
	resLog = zerolog.New(resFile).With().Timestamp().Logger()
	errLog = zerolog.New(errFile).With().Timestamp().Logger()

	suite.Run(t, new(TestTransactionSuite))
}

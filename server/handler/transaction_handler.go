package handler

import (
	"encoding/json"
	"kaspin/server/request"
	"kaspin/server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/thedevsaddam/govalidator"
)

type TransactionHandler struct {
	reqLog zerolog.Logger
	resLog zerolog.Logger
	errLog zerolog.Logger
}

func NewTransactionHandler(reqLog, resLog, errLog zerolog.Logger) TransactionHandler {
	return TransactionHandler{reqLog, resLog, errLog}
}

// Register godoc
// @Summary Register
// @Description Transaction Regist to NICEPAY
// @ID post-register
// @Tags Post Register
// @Accept json
// @Produce json
// @Param params body request.RegisterRequest true "Required data"
// @Success 200
// @Router /register [post]
func (handler TransactionHandler) Register() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get data from request
		var registerRequest request.RegisterRequest
		var err error
		// validation
		rules := govalidator.MapData{
			"amt":            []string{"required"},
			"billingAddr":    []string{"required"},
			"billingCity":    []string{"required"},
			"billingCountry": []string{"required"},
			"billingEmail":   []string{"required"},
			"billingNm":      []string{"required"},
			"billingPhone":   []string{"required"},
			"billingPostCd":  []string{"required"},
			"billingState":   []string{"required"},
			"cartData":       []string{"required"},
			"currency":       []string{"required"},
			"goodsNm":        []string{"required"},
			"instmntMon":     []string{"required"},
			"instmntType":    []string{"required"},
			"payMethod":      []string{"required"},
			"recurrOpt":      []string{"required"},
			"userIP":         []string{"required"},
		}
		opts := govalidator.Options{
			Request: context.Request,
			Data:    &registerRequest,
			Rules:   rules,
		}
		validator := govalidator.New(opts)
		validate := validator.ValidateJSON()
		errValidation := map[string]interface{}{"ErrorValidation": validate}
		if len(validate) != 0 {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusUnprocessableEntity, errValidation)
			return
		}

		reqByte, _ := json.Marshal(registerRequest)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		handler.reqLog.Info().
			RawJSON("payload", reqByte).
			Msg("Register request data (user)")
		srv := service.TransactionService{}
		res, err := srv.Register(registerRequest, handler.reqLog, handler.resLog)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// handle response
		resByte, err := json.Marshal(res)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		handler.resLog.Info().
			RawJSON("payload", resByte).
			Msg("Register response data (user)")
		resultCode, err := strconv.Atoi(res.ResultCd)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if resultCode > 1000 {
			context.JSON(http.StatusBadRequest, res)
			return
		}

		context.JSON(http.StatusOK, res)
		return

	}
}

// Status godoc
// @Summary Status
// @Description Check Status to NICEPAY
// @ID post-status
// @Tags Post Status
// @Accept json
// @Produce json
// @Param params body request.StatusRequest true "Required data"
// @Success 200
// @Router /status [post]
func (handler TransactionHandler) Status() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get data from request
		var statusRequest request.StatusRequest
		var err error
		// validation
		rules := govalidator.MapData{
			"amt":         []string{"required"},
			"tXid":        []string{"required"},
			"referenceNo": []string{"required"},
		}
		opts := govalidator.Options{
			Request: context.Request,
			Data:    &statusRequest,
			Rules:   rules,
		}
		validator := govalidator.New(opts)
		validate := validator.ValidateJSON()
		errValidation := map[string]interface{}{"ErrorValidation": validate}
		if len(validate) != 0 {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusUnprocessableEntity, errValidation)
			return
		}

		reqByte, err := json.Marshal(statusRequest)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		srv := service.TransactionService{}
		handler.reqLog.Info().
			RawJSON("payload", reqByte).
			Msg("Check status request data (user)")
		// process
		res, err := srv.Status(statusRequest, handler.reqLog, handler.resLog)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// handle response
		resByte, err := json.Marshal(res)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		handler.resLog.Info().
			RawJSON("payload", resByte).
			Msg("Check status response data (user)")
		resultCode, err := strconv.Atoi(res.ResultCd)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if resultCode > 1000 {
			context.JSON(http.StatusBadRequest, res)
			return
		}

		context.JSON(http.StatusOK, res)
		return
	}
}

func (handler TransactionHandler) Callback() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get data from request
		var callbackRequest request.CallbackRequest
		var err error
		err = context.ShouldBind(&callbackRequest)
		// validation
		if err != nil {
			context.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		context.JSON(http.StatusOK, callbackRequest)
		return

	}
}

// Payment godoc
// @Summary Payment
// @Description Send Payment to NICEPAY
// @ID post-payment
// @Tags Post Payment
// @Accept json
// @Produce json
// @Param params body request.PaymentRequest true "Required data"
// @Success 200
// @Router /payment [post]
func (handler TransactionHandler) Payment() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get data from request
		var paymentRequest request.PaymentRequest
		var err error
		// validation
		rules := govalidator.MapData{
			"amt":          []string{"required"},
			"cardCvv":      []string{"required"},
			"cardExpYymm":  []string{"required"},
			"cardHolderNm": []string{"required"},
			"cardNo":       []string{"required"},
			"tXid":         []string{"required"},
			"referenceNo":  []string{"required"},
		}
		opts := govalidator.Options{
			Request: context.Request,
			Data:    &paymentRequest,
			Rules:   rules,
		}
		validator := govalidator.New(opts)
		validate := validator.ValidateJSON()
		errValidation := map[string]interface{}{"ErrorValidation": validate}
		if len(validate) != 0 {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusUnprocessableEntity, errValidation)
			return
		}

		reqByte, err := json.Marshal(paymentRequest)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		srv := service.TransactionService{}
		handler.reqLog.Info().
			RawJSON("payload", reqByte).
			Msg("Payment request data (user)")
			// start process
		res, err := srv.Payment(paymentRequest, handler.reqLog, handler.resLog)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// handle response
		resByte, err := json.Marshal(res)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		handler.resLog.Info().
			RawJSON("payload", resByte).
			Msg("Payment response data (user)")
		resultCode, err := strconv.Atoi(res.ResultCd)
		if err != nil {
			handler.errLog.Error().Err(err).Msg("Error Log")
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if resultCode > 1000 {
			context.JSON(http.StatusBadRequest, res)
			return
		}

		context.JSON(http.StatusOK, res)
		return

	}
}

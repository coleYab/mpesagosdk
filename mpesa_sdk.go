package mpesagosdk

import (
	"fmt"
	"net/http"

	"github.com/coleYab/mpesagosdk/account"
	"github.com/coleYab/mpesagosdk/b2c"
	"github.com/coleYab/mpesagosdk/c2b"
	"github.com/coleYab/mpesagosdk/config"
	"github.com/coleYab/mpesagosdk/internal/auth"
	"github.com/coleYab/mpesagosdk/internal/client"
	"github.com/coleYab/mpesagosdk/internal/logger"
	"github.com/coleYab/mpesagosdk/internal/utils"
	"github.com/coleYab/mpesagosdk/transaction"
	"github.com/coleYab/mpesagosdk/types"
	"github.com/go-playground/validator/v10"
)

type App struct {
	cfg       *config.Config
	client    *client.HttpClient
	validator *validator.Validate
    logger    *logger.Logger
}

// Creates a new mpesa App
func New(cfg *config.Config) *App {
	c := client.New(cfg)
	v := validator.New()
    l := logger.NewLogger(logger.ParseLevel(cfg.LogLevel))
	return &App{cfg: cfg, client: c, validator: v, logger: l}
}

func executeRequest[T any](m *App, req types.MpesaRequest, endpoint, method string, authType string) (*T, error) {
	// Validate the request
    masked := utils.MaskEndpoint(endpoint)
    m.logger.Info("making request", "method", method, "endpoint", masked)
	if err := req.Validate(m.validator); err != nil {
        m.logger.Info("validation failed", "error", err.Error())
		return nil, err
	}

	req.FillDefaults()

	response, err := m.client.ApiRequest(m.cfg.Enviroment, endpoint, method, req, authType)
	if err != nil {
        m.logger.Info("request failed", "error", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	// Decode the response and type assert the response failing is impossible
	res, err := req.DecodeResponse(response)
	if err != nil {
        m.logger.Info("request failed", "error", err.Error())
		return nil, err
	}

	resC, ok := res.(T)
	if !ok {
        m.logger.Info("unable to decode the response")
		return nil, fmt.Errorf("unable to decode success message")
	}

    m.logger.Info("request succeded", "method", method, "endpoint", masked)
	return &resC, nil
}

func (m *App) MakeAccountBalanceQuery(req account.AccountBalanceRequest) (*account.AccountBalanceSuccessResponse, error) {
	endpoint := "/mpesa/accountbalance/v1/query"
	return executeRequest[account.AccountBalanceSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *App) MakeB2CPaymentRequest(req b2c.B2CRequest) (*b2c.B2CSuccessResponse, error) {
	endpoint := "/mpesa/b2c/v2/paymentrequest"
	return executeRequest[b2c.B2CSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *App) MakeTransactionReversalRequest(req transaction.TransactionReversalRequest) (*transaction.TransactionReversalResponse, error) {
	endpoint := "/mpesa/reversal/v1/request"
	return executeRequest[transaction.TransactionReversalResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *App) MakeTransactionStatusQuery(req transaction.TransactionStatusRequest) (*transaction.TransactionStatusResponse, error) {
	endpoint := "/mpesa/transactionstatus/v1/query"
	return executeRequest[transaction.TransactionStatusResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *App) USSDPaymentRequest(req c2b.USSDPaymentRequest) (*c2b.USSDSuccessResponse, error) {
	endpoint := "/mpesa/stkpush/v3/processrequest"
	return executeRequest[c2b.USSDSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *App) SimulateCustomerInitiatedPayment(req c2b.SimulateCustomerInititatedPayment) (*c2b.SimulatePaymentSuccessResponse, error) {
	endpoint := "/mpesa/b2c/simulatetransaction/v1/request"
	return executeRequest[c2b.SimulatePaymentSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

func (m *App) RegisterNewURL(req c2b.RegisterC2BURLRequest) (*c2b.RegisterURLResponse, error) {
	endpoint := "/v1/c2b-register-url/register?apikey=" + m.cfg.ConsumerKey
	return executeRequest[c2b.RegisterURLResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeNone)
}

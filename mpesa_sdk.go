package mpesasdk

import (
	"fmt"
	"net/http"

	"github.com/coleYab/mpesasdk/account"
	"github.com/coleYab/mpesasdk/b2c"
	"github.com/coleYab/mpesasdk/c2b"
	"github.com/coleYab/mpesasdk/config"
	"github.com/coleYab/mpesasdk/internal/auth"
	"github.com/coleYab/mpesasdk/internal/client"
	"github.com/coleYab/mpesasdk/transaction"
	"github.com/coleYab/mpesasdk/types"
	"github.com/go-playground/validator/v10"
)

type App struct {
	cfg       *config.Config
	client    *client.HttpClient
	validator *validator.Validate
}

// Creates a new mpesa App
func New(cfg *config.Config) *App {
	c := client.New(cfg)
	v := validator.New()
	return &App{cfg: cfg, client: c, validator: v}
}

func executeRequest[T any](m *App, req types.MpesaRequest, endpoint, method string, authType string) (*T, error) {
	// Validate the request
	if err := req.Validate(m.validator); err != nil {
		return nil, err
	}

	req.FillDefaults()

	response, err := m.client.ApiRequest(m.cfg.Enviroment, endpoint, method, req, authType)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decode the response and type assert the response failing is impossible
	res, err := req.DecodeResponse(response)
	if err != nil {
		return nil, err
	}

	resC, ok := res.(T)
	if !ok {
		return nil, fmt.Errorf("unable to decode success message")
	}

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

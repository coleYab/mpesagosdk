package mpesasdk

import (
	"fmt"
	"net/http"

	"github.com/coleYab/mpesasdk/account"
	"github.com/coleYab/mpesasdk/b2c"
	"github.com/coleYab/mpesasdk/config"
	"github.com/coleYab/mpesasdk/internal/auth"
	"github.com/coleYab/mpesasdk/internal/client"
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

func executeRequest(m *App, req types.MpesaRequest, endpoint, method string, authType string) (types.MpesaResponse, error) {
	// Validate the request
	if err := req.Validate(m.validator); err != nil {
		return nil, err
	}

	req.FillDefaults()

	response, err := m.client.ApiRequest(m.cfg.Enviroment, endpoint, method, req, authType)
	if err != nil {
		fmt.Println("Error: occured here: ", err)
		return nil, err
	}
	defer response.Body.Close()

	// Decode the response and type assert the response failing is impossible
	return req.DecodeResponse(response)
}

func (m *App) MakeAccountBalanceQuery(req account.AccountBalanceRequest) (*account.AccountBalanceSuccessResponse, error) {
	endpoint := "/mpesa/b2c/v2/paymentrequest"
	res, err := executeRequest(m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
	if err != nil {
		return nil, err
	}

	resC, ok := res.(account.AccountBalanceSuccessResponse)
	if !ok {
		return nil, fmt.Errorf("unable to decode success message")
	}

	return &resC, nil
}

func (m *App) MakeB2CPaymentRequest(req b2c.B2CRequest) (*b2c.B2CSuccessResponse, error) {
	endpoint := "/mpesa/b2c/v2/paymentrequest"
	res, err := executeRequest(m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
	if err != nil {
		return nil, err
	}

	resC, ok := res.(b2c.B2CSuccessResponse)
	if !ok {
		return nil, fmt.Errorf("unable to decode success message")
	}

	return &resC, nil
}

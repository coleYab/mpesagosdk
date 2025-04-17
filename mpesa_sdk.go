// Package mpesagosdk provides a comprehensive SDK for interacting with the M-Pesa API,
// allowing developers to integrate mobile money payment services into their applications.
// The package simplifies common M-Pesa operations such as C2B (Customer-to-Business) payments,
// B2C (Business-to-Customer) disbursements, transaction status checks, account balance inquiries,
// and transaction reversals.
//
// The package is designed to abstract the complexities of working directly with M-Pesa APIs by
// providing strongly-typed request/response structures, reusable client configurations, and automatic
// request validation and default value population.
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

// App: represents an instance of the M-Pesa application SDK.
// This struct serves as the main entry point for interacting with the M-Pesa API,
// providing methods to initiate various types of M-Pesa transactions, such as B2C payments,
// account balance queries, transaction status queries, and more.
//
// Fields:
//	- `cfg`: The configuration settings for the SDK, containing credentials (consumer key/secret),
//     environment settings (SANDBOX or PRODUCTION), log level, etc.
//	- `client`: The HTTP client used to make requests to the M-Pesa API.
//	- `validator`: A validator instance that ensures requests conform to the expected structure.
//	- `logger`: A logger instance that logs relevant information and errors for monitoring and debugging.
//
// Example Usage:
//
//	package main
//
//
//	...(your code goes here)
//
//	cfg, err := config.NewFromEnv()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Create a new instance of the App struct
//	app := mpesagosdk.New(cfg)
//
//	// Now you can use the app to interact with M-Pesa API, for example:
//	res, err := app.MakeB2CPaymentRequest(someB2CPaymentRequest)
//	// Handle response or error
//	...
type App struct {
	cfg       *config.Config
	client    *client.HttpClient
	validator *validator.Validate
	logger    *logger.Logger
}

// New: Creates a new instance of the M-Pesa App.
// The New function initializes a new App instance by accepting a Config struct, which contains the necessary configuration data such as the ConsumerKey, ConsumerSecret, LogLevel, and environment settings. It then creates and returns an instance of App with all necessary components, such as:
//
//	- A newly initialized HTTP client (client.HttpClient) to interact with the M-Pesa API.
//	- A newly initialized validator.Validate instance for validating requests.
//	- A logger (logger.Logger) to track the SDK's activities at the specified log level.
//
// Parameters:
//	- cfg: A pointer to a config.Config struct containing the required configuration data for the M-Pesa SDK.
//
// Returns:
//	- A pointer to an App instance, which is ready to make requests to the M-Pesa API.
func New(cfg *config.Config) *App {
	c := client.New(cfg)
	v := validator.New()
	l := logger.NewLogger(logger.ParseLevel(cfg.LogLevel))
	return &App{cfg: cfg, client: c, validator: v, logger: l}
}

// executeRequest: this is a function that will act as an orchestrator that can help us in determinig what
// to do with the request. It has three main steps.
//	1. Validation: here it will use the validation is defined at the types.MpesaRequest struct
// 	2. FillDefault: it will fill the default data that is unique and default to each request
// 	3. Api request: then it will utilize the clients ability to send ApiRequest and sends an api request
// 	4. DecodeResponse: finally it will decode the response that comes from the mpesa.
//
// Returns:
// 	- T: generic type that has to be specified on success
// 	- error: on failure.
func executeRequest[T any](m *App, req types.MpesaRequest, endpoint, method string, authType string) (*T, error) {
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

// MakeAccountBalanceQuery: Sends a request to the M-Pesa API to query the balance
// of an account. It uses the `AccountBalanceRequest` struct to specify the necessary
// parameters and returns the response in the form of `AccountBalanceSuccessResponse`.
//
// This function performs a POST request to the "/mpesa/accountbalance/v1/query" endpoint
// using the Bearer token for authentication.
//
// Parameters:
//	- `req`: The `AccountBalanceRequest` struct containing the necessary request data
//     for querying the account balance.
//
// Returns:
//	- A pointer to the `AccountBalanceSuccessResponse` struct on a successful request.
//	- An error if the request fails or is invalid.
func (m *App) MakeAccountBalanceQuery(req account.AccountBalanceRequest) (*account.AccountBalanceSuccessResponse, error) {
	endpoint := "/mpesa/accountbalance/v1/query"
	return executeRequest[account.AccountBalanceSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

// MakeB2CPaymentRequest: Sends a request to the M-Pesa API to initiate a B2C (Business
// to Customer) payment. It uses the `B2CRequest` struct to specify the payment details
// and returns the response in the form of `B2CSuccessResponse`.
//
// This function performs a POST request to the "/mpesa/b2c/v2/paymentrequest" endpoint
// using the Bearer token for authentication.
//
// Parameters:
//	- `req`: The `B2CRequest` struct containing the necessary request data for initiating
//     a B2C payment.
//
// Returns:
//	- A pointer to the `B2CSuccessResponse` struct on a successful request.
//	- An error if the request fails or is invalid.
func (m *App) MakeB2CPaymentRequest(req b2c.B2CRequest) (*b2c.B2CSuccessResponse, error) {
	endpoint := "/mpesa/b2c/v2/paymentrequest"
	return executeRequest[b2c.B2CSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

// MakeTransactionReversalRequest: Sends a request to the M-Pesa API to reverse a previously
// processed transaction. It uses the `TransactionReversalRequest` struct to specify
// the transaction details to be reversed and returns the response in the form of
// `TransactionReversalResponse`.
//
// This function performs a POST request to the "/mpesa/reversal/v1/request" endpoint
// using the Bearer token for authentication.
//
// Parameters:
//	- `req`: The `TransactionReversalRequest` struct containing the necessary request data
//     for reversing a transaction.
//
// Returns:
//	- A pointer to the `TransactionReversalResponse` struct on a successful request.
//	- An error if the request fails or is invalid.
func (m *App) MakeTransactionReversalRequest(req transaction.TransactionReversalRequest) (*transaction.TransactionReversalResponse, error) {
	endpoint := "/mpesa/reversal/v1/request"
	return executeRequest[transaction.TransactionReversalResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

// MakeTransactionStatusQuery: Sends a request to the M-Pesa API to check the status of
// a specific transaction. It uses the `TransactionStatusRequest` struct to specify the
// transaction details and returns the response in the form of `TransactionStatusResponse`.
//
// This function performs a POST request to the "/mpesa/transactionstatus/v1/query" endpoint
// using the Bearer token for authentication.
//
// Parameters:
//	- `req`: The `TransactionStatusRequest` struct containing the necessary request data
//     for querying the status of a transaction.
//
// Returns:
// 	- A pointer to the `TransactionStatusResponse` struct on a successful request.
// 	- An error if the request fails or is invalid.
func (m *App) MakeTransactionStatusQuery(req transaction.TransactionStatusRequest) (*transaction.TransactionStatusResponse, error) {
	endpoint := "/mpesa/transactionstatus/v1/query"
	return executeRequest[transaction.TransactionStatusResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

// USSDPaymentRequest: Sends a request to the M-Pesa API to initiate a USSD (Unstructured
// Supplementary Service Data) payment. It uses the `USSDPaymentRequest` struct to specify
// the payment details and returns the response in the form of `USSDSuccessResponse`.
//
// This function performs a POST request to the "/mpesa/stkpush/v3/processrequest" endpoint
// using the Bearer token for authentication.
//
// Parameters:
// 	- `req`: The `USSDPaymentRequest` struct containing the necessary request data for initiating
//     a USSD payment.
//
// Returns:
// 	- A pointer to the `USSDSuccessResponse` struct on a successful request.
// 	- An error if the request fails or is invalid
func (m *App) USSDPaymentRequest(req c2b.USSDPaymentRequest) (*c2b.USSDSuccessResponse, error) {
	endpoint := "/mpesa/stkpush/v3/processrequest"
	return executeRequest[c2b.USSDSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

// SimulateCustomerInitiatedPayment: Sends a request to the M-Pesa API to simulate
// a customer-initiated payment. It uses the `SimulateCustomerInititatedPayment` struct
// to specify the payment details and returns the response in the form of
// `SimulatePaymentSuccessResponse`.
//
// This function performs a POST request to the "/mpesa/b2c/simulatetransaction/v1/request"
// endpoint using the Bearer token for authentication.
//
// Parameters:
// 	- `req`: The `SimulateCustomerInititatedPayment` struct containing the necessary request
//     data for simulating a customer-initiated payment.
//
// Warning:
// 	- This endpoint may generate some weird behaviours due to unstablity please contact the
// mpesa developer team to fix this issue if you really need that feature out there.
//
// Returns:
// 	- A pointer to the `SimulatePaymentSuccessResponse` struct on a successful request.
// 	- An error if the request fails or is invalid.
func (m *App) SimulateCustomerInitiatedPayment(req c2b.SimulateCustomerInititatedPayment) (*c2b.SimulatePaymentSuccessResponse, error) {
	endpoint := "/mpesa/b2c/simulatetransaction/v1/request"
	return executeRequest[c2b.SimulatePaymentSuccessResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeBearer)
}

// RegisterNewURL: Sends a request to the M-Pesa API to register a new callback URL for
// C2B (Customer to Business) transactions. It uses the `RegisterC2BURLRequest` struct
// to specify the URL to be registered and returns the response in the form of
// `RegisterURLResponse`.
//
// This function performs a POST request to the "/v1/c2b-register-url/register" endpoint
// using no authentication token (AuthTypeNone) and includes the `ConsumerKey` as part of the query string.
//
// Parameters:
//	- `req`: The `RegisterC2BURLRequest` struct containing the necessary request data for
//     registering a new URL.
//
// Returns:
//	- A pointer to the `RegisterURLResponse` struct on a successful request.
//	- An error if the request fails or is invalid.
func (m *App) RegisterNewURL(req c2b.RegisterC2BURLRequest) (*c2b.RegisterURLResponse, error) {
	endpoint := "/v1/c2b-register-url/register?apikey=" + m.cfg.ConsumerKey
	return executeRequest[c2b.RegisterURLResponse](m, &req, endpoint, http.MethodPost, auth.AuthTypeNone)
}

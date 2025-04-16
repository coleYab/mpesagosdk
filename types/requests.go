// Package types defines the structures and interfaces for working with M-Pesa API requests and responses.
// It includes interfaces for validating, decoding, and handling M-Pesa API request and response data.
// The package also defines common response structures, error handling, and success responses.
//
// Key Features:
// 	- Defines the `MpesaRequest` interface to validate, decode, and fill defaults for requests.
// 	- Defines the `MpesaResponse` interface for all successful responses.
// 	- Provides structures for handling error responses from M-Pesa APIs (`MpesaErrorResponse`).
// 	- Defines a common response structure (`MpesaCommonResponse`) with shared fields.
// 	- Provides an error handling mechanism for failed requests via the `Error` method on `MpesaErrorResponse`.
package types

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// MpesaRequest defines the interface that must be implemented by all request types
// to interact with the M-Pesa API.
//
// This interface ensures that all request types provide standard methods for:
//   - Decoding responses from the M-Pesa API.
//   - Validating the request data before sending it to the API.
//   - Populating default values for some fields with known value.
//
// Implementing this interface allows request types to decode success response
// seamlessly with the M-Pesa API client. This features insures consistency when handing
// the addition of new endpoint to mpesa api.
//
// Methods:
//
//   DecodeResponse(res *http.Response) (MpesaResponse, error):
//     Decodes the HTTP response from the M-Pesa API into an appropriate Go structure.
//
//   Validate() error:
//     Performs validation checks on the request fields to ensure the data is complete and correct
//     before sending the request to the API. Returns an error if the validation fails.
//
//   FillDefaults():
//     Populates default values for fields in the request. This ensures required fields
//     have valid defaults if not explicitly set by the user.
//
// Example:
//   To create a new request type for an M-Pesa API endpoint, define a struct for the request data
//   and implement the MpesaRequest interface. For example if mpesa decided to add a new feature
//   for collecting recurring payments:
//
//   ```go
//   // add the required json and validation tags
//   type RecurringPaymentRequest struct {
//       Field1 string `json:"Field1"`
//       Field2 int    `json:"Field2"`
// 		 ...
//   }
//
//   func (r *MyMpesaRequest) DecodeResponse(res *http.Response) (interface{}, error) {
//       // Implement response decoding logic
//   }
//
//   func (r *MyMpesaRequest) Validate() error {
//       // Implement validation logic
//   }
//
//   func (r *MyMpesaRequest) FillDefaults() {
//       // Set default values for fields
//   }
//   ```
//
// This interface will improve consistency when adding a new feature to the sandbox.
type MpesaRequest interface {
	Validate(v *validator.Validate) error
	DecodeResponse(res *http.Response) (MpesaResponse, error)
	FillDefaults()
}

// MpesaResponse: this response should be returned only on success
type MpesaResponse interface{}

// MpesaErrorResponse is a structure used to represent error responses from the M-Pesa API.
// It contains information about the error, including the request ID, error code, and error message.
// it implements the (error interface)[https://go.dev/wiki/Errors]
type MpesaErrorResponse struct {
	RequestId    string `json:"requestId"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

// TODO: handle the error response please do that efficently
func (e *MpesaErrorResponse) Error() string {
	return fmt.Sprintf("request with id=%v failed with code=%v, due to %v", e.RequestId, e.ErrorCode, e.ErrorMessage)
}

type MpesaSuccessResponse struct{}


// MpesaCommonResponse is a structure that holds fields common to all M-Pesa API responses.
// It includes the conversation ID, originator conversation ID, response description, and response code.
type MpesaCommonResponse struct {
	RequestType             string
	ConversationID          string `json:"ConversationID"`
	OriginatorConversatonId string `json:"OriginatorConversationID"`
	ResponseDescription     string `json:"ResponseDescription"`
	ResponseCode            string `json:"ResponseCode"`
}

package types

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type MpesaRequest interface {
    Validate(v *validator.Validate) error
    DecodeResponse(res *http.Response) (MpesaResponse, error)
    FillDefaults()
}

// MpesaResponse: this response should be returned only on success
type MpesaResponse interface {}

type MpesaErrorResponse struct {
    RequestId    string `json:"requestId"`
    ErrorCode    string `json:"errorCode"`
    ErrorMessage string `json:"errorMessage"`
}

type MpesaSuccessResponse struct { }

// TODO: handle the error response please do that efficently
func (e *MpesaErrorResponse) Error() string {
    return fmt.Sprintf("request with id=%v failed with code=%v, due to %v", e.RequestId, e.ErrorCode, e.ErrorMessage)
}

type MpesaCommonResponse struct {
    RequestType             string
    ConversationID          string `json:"ConversationID"`
    OriginatorConversatonId string `json:"OriginatorConversationID"`
    ResponseDescription     string `json:"ResponseDescription"`
    ResponseCode            string `json:"ResponseCode"`
}

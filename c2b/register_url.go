package c2b

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/coleYab/mpesasdk/types"
)


type RegisterC2BURLRequest struct {
	ShortCode string `json:"ShortCode" validate:"required,min=5,max=20"`
	ResponseType types.ResponseType `json:"ResponseType" validate:"required"`
	CommandID types.CommandId `json:"CommandID" validate:"required"`
	ConfirmationURL string `json:"ConfirmationURL" validate:"required,url"`
	ValidationURL   string `json:"ValidationURL" validate:"required,url"`
}

type registerUrlResponse struct {
	Header struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		CustomerMessage string `json:"customerMessage"`
	} `json:"header"`
}

type RegisterC2BURLSuccessResponse types.MpesaCommonResponse

func (s *RegisterC2BURLRequest) DecodeResponse(res *http.Response) (types.MpesaResponse, error) {
	bodyData, _ := io.ReadAll(res.Body)
	responseData := registerUrlResponse{}
	err := json.Unmarshal(bodyData, &responseData)
	if err != nil {
        return nil, err
	}

	switch responseData.Header.ResponseCode {
	case "200":
		return RegisterC2BURLSuccessResponse{
			ResponseCode:        string(responseData.Header.ResponseCode),
			ResponseDescription: responseData.Header.ResponseMessage,
		}, nil
	case "":
		errorResponse := types.MpesaErrorResponse{}
		err := json.Unmarshal(bodyData, &errorResponse)
		if err != nil {
            return nil, err
		}
		responseData.Header.ResponseCode = errorResponse.ErrorCode
		responseData.Header.ResponseMessage = errorResponse.ErrorMessage
	}

    return nil, &types.MpesaErrorResponse{
        ErrorCode: responseData.Header.ResponseCode,
        ErrorMessage: responseData.Header.ResponseMessage,
    }
}

func (t *RegisterC2BURLRequest) FillDefaults() {
	t.CommandID = types.RegisterURLCommand
}

func (t *RegisterC2BURLRequest) Validate() error {

	validResponseType := []types.ResponseType{types.CompletedResponse, types.CancelledResponse}
	if !slices.Contains(validResponseType, t.ResponseType) {
        return fmt.Errorf("unkown response type: %v", t.ResponseType)
    }

	return nil
}

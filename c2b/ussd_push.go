package c2b

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/coleYab/mpesasdk/internal/utils"
	"github.com/coleYab/mpesasdk/types"
	"github.com/go-playground/validator/v10"
)

type ReferenceDataRequest struct {
	Key   string `json:"Key" validate:"required,min=1,max=50"`
	Value string `json:"Value" validate:"required,min=1,max=100"`
}

type USSDPaymentRequest struct {
	MerchantRequestID string `json:"MerchantRequestID" validate:"required,min=1,max=50"`

	BusinessShortCode string `json:"BusinessShortCode" validate:"required"`

	ReferenceData []ReferenceDataRequest `json:"ReferenceData" validate:"dive"`

	TransactionType types.TransactionType `json:"TransactionType" validate:"required"`

	Password string `json:"Password" validate:"required,min=8,max=100"`

	Timestamp string `json:"Timestamp" validate:"required,datetime=20060102150405"`

	Amount uint64 `json:"Amount" validate:"required,gt=0"`

	PartyA string `json:"PartyA" validate:"required,min=1,max=20"`

	PartyB string `json:"PartyB" validate:"required,min=1,max=20"`

	PhoneNumber string `json:"PhoneNumber" validate:"required"`

	CallBackURL string `json:"CallBackURL" validate:"required,url"`

	AccountReference string `json:"AccountReference" validate:"required,min=1,max=20"`

	TransactionDesc string `json:"TransactionDesc" validate:"required,min=1,max=100"`
}

type USSDSuccessResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	CustomerMessage     string `json:"CustomerMessage"`
	ResponseDescription string `json:"ResponseDescription"`
}

type USSDRequestError USSDSuccessResponse

func (s *USSDPaymentRequest) DecodeResponse(res *http.Response) (types.MpesaResponse, error) {
	bodyData, _ := io.ReadAll(res.Body)
	responseData := USSDSuccessResponse{}
	err := json.Unmarshal(bodyData, &responseData)
	if err != nil {
		return nil, err
	}

	switch responseData.ResponseCode {
	case "0":
		return &responseData, nil
	case "":
		e := types.MpesaErrorResponse{}
		err := json.Unmarshal(bodyData, &e)
		if err != nil {
			return nil, err
		}
	}

	return nil, &types.MpesaErrorResponse{
		RequestId:    responseData.MerchantRequestID,
		ErrorCode:    responseData.ResponseCode,
		ErrorMessage: responseData.ResponseDescription,
	}
}

func (t *USSDPaymentRequest) FillDefaults() {
}

func (t *USSDPaymentRequest) Validate(v *validator.Validate) error {
	validTransactionTypes := []types.TransactionType{
		types.CustomerBuyGoodsOnlineTransaction,
		types.CustomerPayBillOnlineTransaction,
	}

	if !slices.Contains(validTransactionTypes, t.TransactionType) {
		return fmt.Errorf("invalid transaction type: %v", t.TransactionType)
	}

	return utils.Validate(v, t)
}

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

type SimulateCustomerInititatedPayment struct {
	CommandID     types.CommandId `json:"CommandID" validate:"required"`
	Amount        uint64          `json:"Amount" validate:"required,gte=1"`
	Msisdn        string          `json:"Msisdn" validate:"required,numeric"`
	BillRefNumber string          `json:"BillRefNumber" validate:"required,min=6,max=20"`
	ShortCode     string          `json:"ShortCode" validate:"required"`
}

type SimulatePaymentSuccessResponse types.MpesaCommonResponse

func (s *SimulateCustomerInititatedPayment) DecodeResponse(res *http.Response) (types.MpesaResponse, error) {
	bodyData, _ := io.ReadAll(res.Body)
	responseData := SimulatePaymentSuccessResponse{}
	err := json.Unmarshal(bodyData, &responseData)
	if err != nil {
		return nil, err
	}

	if responseData.ResponseCode != "0" {
		errorResponseData := types.MpesaErrorResponse{}
		err := json.Unmarshal(bodyData, &errorResponseData)
		if err != nil {
			return SimulatePaymentSuccessResponse{}, err
		}
		return nil, &errorResponseData
	}

	return responseData, nil
}

func (s *SimulateCustomerInititatedPayment) FillDefaults() {}

func (s *SimulateCustomerInititatedPayment) Validate(v *validator.Validate) error {
	validCommands := []types.CommandId{
		types.CustomerPayBillOnlineCommand, types.CustomerBuyGoodsOnlineCommand,
	}
	if !slices.Contains(validCommands, s.CommandID) {
		return fmt.Errorf("unkown command %v", s.CommandID)
	}

	return utils.Validate(v, s)
}

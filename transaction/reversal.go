package transaction

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/coleYab/mpesagosdk/internal/utils"
	"github.com/coleYab/mpesagosdk/types"
	"github.com/go-playground/validator/v10"
)

type TransactionReversalRequest struct {
	Initiator                string               `json:"Initiator" validate:"required,min=3,max=50"`
	SecurityCredential       string               `json:"SecurityCredential" validate:"required"`
	CommandID                types.CommandId      `json:"CommandID" validate:"required"`
	TransactionID            string               `json:"TransactionID" validate:"required,min=10,max=100"`
	Amount                   uint64               `json:"Amount" validate:"required,gte=1"`
	PartyA                   string               `json:"ReceiverParty" validate:"required,min=3,max=50"`
	IdentifierType           types.IdentifierType `json:"RecieverIdentifierType" validate:"required"`
	QueueTimeOutURL          string               `json:"QueueTimeOutURL" validate:"required,url"`
	ResultURL                string               `json:"ResultURL" validate:"required,url"`
	Remarks                  string               `json:"Remarks" validate:"omitempty,max=200"`
	Occasion                 string               `json:"Occasion" validate:"omitempty,max=100"`
    OriginatorConversationID string               `json:"OriginatorConversationID" validate:"required"`
}

type TransactionReversalResponse types.MpesaCommonResponse

func (a *TransactionReversalRequest) DecodeResponse(res *http.Response) (types.MpesaResponse, error) {
	bodyData, _ := io.ReadAll(res.Body)
	responseData := TransactionReversalResponse{}
	err := json.Unmarshal(bodyData, &responseData)
	if err != nil {
		return nil, err
	}

	if responseData.ResponseCode != "0" {
		errorResponseData := types.MpesaErrorResponse{}
		err := json.Unmarshal(bodyData, &errorResponseData)
		if err != nil {
			return nil, err
		}

		return nil, &errorResponseData
	}

	return responseData, nil
}

func (a *TransactionReversalRequest) FillDefaults() {
	a.CommandID = types.TransactionReversalCommand
}

func (a *TransactionReversalRequest) Validate(v *validator.Validate) error {
	return utils.Validate(v, a)
}

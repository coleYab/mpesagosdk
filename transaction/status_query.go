package transaction

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

// TransactionStatusRequest represents the parameters for querying the transaction status.
type TransactionStatusRequest struct {
	CommandID                types.CommandId      `json:"CommandID" validate:"required"`
	IdentifierType           types.IdentifierType `json:"IdentifierType" validate:"required"`
	Initiator                string               `json:"Initiator" validate:"required,min=1,max=255"`
	Occasion                 string               `json:"Occasion" validate:"required,min=1,max=255"`
	OriginatorConversationID string               `json:"OriginatorConversationID,omitempty" validate:"omitempty,min=1,max=255"`
	PartyA                   string               `json:"PartyA" validate:"required,min=1,max=255"`
	QueueTimeOutURL          string               `json:"QueueTimeOutURL" validate:"required,url"`
	Remarks                  string               `json:"Remarks" validate:"omitempty,max=500"`
	ResultURL                string               `json:"ResultURL" validate:"required,url"`
	SecurityCredential       string               `json:"SecurityCredential" validate:"required,min=8"`
	TransactionID            string               `json:"TransactionID" validate:"required,min=1,max=255"`
}

type TransactionStatusResponse types.MpesaCommonResponse

func (a *TransactionStatusRequest) DecodeResponse(res *http.Response) (types.MpesaResponse, error) {
	bodyData, _ := io.ReadAll(res.Body)
	responseData := TransactionStatusResponse{}
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

func (a *TransactionStatusRequest) FillDefaults() {
	a.CommandID = types.TransactionStatusCommand
}

func (a *TransactionStatusRequest) Validate(v *validator.Validate) error {
	validIdentifiers := []types.IdentifierType{types.MsisdnIdentifierType, types.TillNumberIdentifierType, types.ShortCodeIdentifierType}
	if !slices.Contains(validIdentifiers, a.IdentifierType) {
		return fmt.Errorf("unkown identifier type %v", a.IdentifierType)
	}

	return utils.Validate(v, a)
}

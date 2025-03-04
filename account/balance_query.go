package account

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/coleYab/mpesasdk/types"
	"github.com/go-playground/validator/v10"
)

type AccountBalanceRequest struct {
    CommandID types.CommandId `json:"CommandID" validate:"required"`

    IdentifierType types.IdentifierType `json:"IdentifierType" validate:"required"`

    Initiator string `json:"Initiator" validate:"required,min=1,max=255"`

    PartyA int `json:"PartyA" validate:"required,min=1"`

    QueueTimeOutURL string `json:"QueueTimeOutURL" validate:"required,url"`

    Remarks string `json:"Remarks" validate:"max=500"`

    ResultURL string `json:"ResultURL" validate:"required,url"`

    SecurityCredential string `json:"SecurityCredential" validate:"required,min=8"`

    OriginatorConversationID string `json:"OriginatorConversationID" validate:"required"`
}

type AccountBalanceSuccessResponse types.MpesaCommonResponse

func (a *AccountBalanceRequest) DecodeResponse(res *http.Response) (types.MpesaResponse, error) {
    bodyData, _ := io.ReadAll(res.Body)
    responseData := AccountBalanceSuccessResponse{}
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

func (a *AccountBalanceRequest) FillDefaults() {
    a.CommandID = types.AccountBalanceCommand
}

func (a *AccountBalanceRequest) Validate(v *validator.Validate) error {
	if err := v.Struct(a); err != nil {
		errCasted := err.(validator.ValidationErrors)
		return errCasted
	}

    validIdentifiers := []types.IdentifierType{types.MsisdnIdentifierType, types.TillNumberIdentifierType, types.ShortCodeIdentifierType}
    if !slices.Contains(validIdentifiers, a.IdentifierType) {
        return fmt.Errorf("unkown identifier type %v", a.IdentifierType)
    }

    return nil
}


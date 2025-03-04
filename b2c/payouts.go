package b2c

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/coleYab/mpesasdk/types"
	"github.com/go-playground/validator/v10"
)

type B2CRequest struct {
	InitiatorName            string          `json:"InitiatorName" validate:"required"`
	SecurityCredential       string          `json:"SecurityCredential" validate:"required"`
	CommandID                types.CommandId `json:"CommandID" validate:"required"`
	Amount                   uint            `json:"Amount" validate:"required,gte=1"`
	PartyA                   uint            `json:"PartyA" validate:"required"`
	PartyB                   uint            `json:"PartyB" validate:"required"`
	Remarks                  string          `json:"Remarks" validate:"required,max=200"`
	QueueTimeOutURL          string          `json:"QueueTimeOutURL" validate:"required,url"`
	ResultURL                string          `json:"ResultURL" validate:"required,url"`
	Occasion                 string          `json:"Occasion" validate:"required"`
	OriginatorConversationID string          `json:"OriginatorConversationID" validate:"required"`
}

// DecodeResponse implements types.MpesaRequest.
func (b *B2CRequest) DecodeResponse(res *http.Response) (types.MpesaResponse, error) {
	bodyData, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	responseData := B2CSuccessResponse{}
	err := json.Unmarshal(bodyData, &responseData)
	fmt.Println("Printing the response: ", responseData)
	fmt.Println(string(bodyData))
	if err != nil {
		return nil, err
	}

	if responseData.ResponseCode != "0" {
		errorResponseData := &types.MpesaErrorResponse{}
		err := json.Unmarshal(bodyData, &errorResponseData)
		if err != nil {
			return nil, err
		}

		return nil, errorResponseData
	}

	return responseData, nil
}

// FillDefaults implements types.MpesaRequest.
func (b *B2CRequest) FillDefaults() {}

// Validate implements types.MpesaRequest.
func (b *B2CRequest) Validate(v *validator.Validate) error {
	if err := v.Struct(b); err != nil {
		errCasted := err.(validator.ValidationErrors)
		return errCasted
	}

	// TODO: check the command id
	// validCommands := []types.CommandId{}
	// if !slices.Contains(validCommands, b.CommandID) {
	// 	return fmt.Errorf("unknown CommandID %v", string(b.CommandID))
	// }

	return nil
}

type B2CSuccessResponse types.MpesaCommonResponse

// TODO: Better error handling

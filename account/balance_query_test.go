package account

import (
	"testing"

	"github.com/coleYab/mpesasdk/types"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestAccountBalanceRequestValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		req     AccountBalanceRequest
		wantErr bool
	}{
		{
			name: "Valid Input",
			req: AccountBalanceRequest{
				CommandID:                  types.CommandId("validCommandID"),
				IdentifierType:             types.IdentifierType("validIdentifierType"),
				Initiator:                  "validInitiator",
				PartyA:                     12345,
				QueueTimeOutURL:            "https://example.com/timeout",
				Remarks:                    "Valid remarks",
				ResultURL:                  "https://example.com/result",
				SecurityCredential:         "validSecurityCredential",
				OriginatorConversationID:   "validOriginatorConversationID",
			},
			wantErr: false,
		},
		{
			name: "Missing CommandID",
			req: AccountBalanceRequest{
				IdentifierType:             types.IdentifierType("validIdentifierType"),
				Initiator:                  "validInitiator",
				PartyA:                     12345,
				QueueTimeOutURL:            "https://example.com/timeout",
				Remarks:                    "Valid remarks",
				ResultURL:                  "https://example.com/result",
				SecurityCredential:         "validSecurityCredential",
				OriginatorConversationID:   "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Invalid IdentifierType",
			req: AccountBalanceRequest{
				CommandID:                  types.CommandId("validCommandID"),
				IdentifierType:             types.IdentifierType("invalidIdentifierType"),
				Initiator:                  "validInitiator",
				PartyA:                     12345,
				QueueTimeOutURL:            "https://example.com/timeout",
				Remarks:                    "Valid remarks",
				ResultURL:                  "https://example.com/result",
				SecurityCredential:         "validSecurityCredential",
				OriginatorConversationID:   "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Initiator Too Short",
			req: AccountBalanceRequest{
				CommandID:                  types.CommandId("validCommandID"),
				IdentifierType:             types.IdentifierType("validIdentifierType"),
				Initiator:                  "", // Empty string
				PartyA:                     12345,
				QueueTimeOutURL:            "https://example.com/timeout",
				Remarks:                    "Valid remarks",
				ResultURL:                  "https://example.com/result",
				SecurityCredential:         "validSecurityCredential",
				OriginatorConversationID:   "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "PartyA Less Than 1",
			req: AccountBalanceRequest{
				CommandID:                  types.CommandId("validCommandID"),
				IdentifierType:             types.IdentifierType("validIdentifierType"),
				Initiator:                  "validInitiator",
				PartyA:                     0, // Invalid value
				QueueTimeOutURL:            "https://example.com/timeout",
				Remarks:                    "Valid remarks",
				ResultURL:                  "https://example.com/result",
				SecurityCredential:         "validSecurityCredential",
				OriginatorConversationID:   "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Invalid QueueTimeOutURL",
			req: AccountBalanceRequest{
				CommandID:                  types.CommandId("validCommandID"),
				IdentifierType:             types.IdentifierType("validIdentifierType"),
				Initiator:                  "validInitiator",
				PartyA:                     12345,
				QueueTimeOutURL:            "invalid-url", // Invalid URL
				Remarks:                    "Valid remarks",
				ResultURL:                  "https://example.com/result",
				SecurityCredential:         "validSecurityCredential",
				OriginatorConversationID:   "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Invalid ResultURL",
			req: AccountBalanceRequest{
				CommandID:                  types.CommandId("validCommandID"),
				IdentifierType:             types.IdentifierType("validIdentifierType"),
				Initiator:                  "validInitiator",
				PartyA:                     12345,
				QueueTimeOutURL:            "https://example.com/timeout",
				Remarks:                    "Valid remarks",
				ResultURL:                  "invalid-url", // Invalid URL
				SecurityCredential:         "validSecurityCredential",
				OriginatorConversationID:   "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "SecurityCredential Too Short",
			req: AccountBalanceRequest{
				CommandID:                  types.CommandId("validCommandID"),
				IdentifierType:             types.IdentifierType("validIdentifierType"),
				Initiator:                  "validInitiator",
				PartyA:                     12345,
				QueueTimeOutURL:            "https://example.com/timeout",
				Remarks:                    "Valid remarks",
				ResultURL:                  "https://example.com/result",
				SecurityCredential:         "short", // Less than 8 characters
				OriginatorConversationID:   "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Remarks Too Long",
			req: AccountBalanceRequest{
				CommandID:                  types.CommandId("validCommandID"),
				IdentifierType:             types.IdentifierType("validIdentifierType"),
				Initiator:                  "validInitiator",
				PartyA:                     12345,
				QueueTimeOutURL:            "https://example.com/timeout",
				Remarks:                    string(make([]byte, 501)), // 501 characters
				ResultURL:                  "https://example.com/result",
				SecurityCredential:         "validSecurityCredential",
				OriginatorConversationID:   "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Missing OriginatorConversationID",
			req: AccountBalanceRequest{
				CommandID:                  types.CommandId("validCommandID"),
				IdentifierType:             types.IdentifierType("validIdentifierType"),
				Initiator:                  "validInitiator",
				PartyA:                     12345,
				QueueTimeOutURL:            "https://example.com/timeout",
				Remarks:                    "Valid remarks",
				ResultURL:                  "https://example.com/result",
				SecurityCredential:         "validSecurityCredential",
				OriginatorConversationID:   "", // Missing
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

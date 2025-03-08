package account

import (
	"testing"

	"github.com/coleYab/mpesasdk/types"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestAccountBalanceRequestValidation(t *testing.T) {
    v := validator.New()
	tests := []struct {
		name    string
		req     AccountBalanceRequest
		wantErr bool
	}{
		{
			name: "Valid Input",
			req: AccountBalanceRequest{
				CommandID:                types.AccountBalanceCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				Initiator:                "apiuser",
				PartyA:                   600000,
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				Remarks:                  "Salary Payment",
				ResultURL:                "https://yourdomain.com/result",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "validOriginatorConversationID",
			},
			wantErr: false,
		},
		{
			name: "Missing CommandID",
			req: AccountBalanceRequest{
				IdentifierType:           types.ShortCodeIdentifierType,
				Initiator:                "apiuser",
				PartyA:                   600000,
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				Remarks:                  "Salary Payment",
				ResultURL:                "https://yourdomain.com/result",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Invalid IdentifierType",
			req: AccountBalanceRequest{
				CommandID:                types.AccountBalanceCommand,
				IdentifierType:           types.IdentifierType("invalidIdentifierType"),
				Initiator:                "apiuser",
				PartyA:                   600000,
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				Remarks:                  "Salary Payment",
				ResultURL:                "https://yourdomain.com/result",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Initiator Too Short",
			req: AccountBalanceRequest{
				CommandID:                types.AccountBalanceCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				Initiator:                "", // Empty string
				PartyA:                   600000,
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				Remarks:                  "Salary Payment",
				ResultURL:                "https://yourdomain.com/result",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "PartyA Less Than 1",
			req: AccountBalanceRequest{
				CommandID:                types.AccountBalanceCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				Initiator:                "apiuser",
				PartyA:                   0, // Invalid value
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				Remarks:                  "Salary Payment",
				ResultURL:                "https://yourdomain.com/result",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Invalid QueueTimeOutURL",
			req: AccountBalanceRequest{
				CommandID:                types.AccountBalanceCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				Initiator:                "apiuser",
				PartyA:                   600000,
				QueueTimeOutURL:          "invalid-url", // Invalid URL
				Remarks:                  "Salary Payment",
				ResultURL:                "https://yourdomain.com/result",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Invalid ResultURL",
			req: AccountBalanceRequest{
				CommandID:                types.AccountBalanceCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				Initiator:                "apiuser",
				PartyA:                   600000,
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				Remarks:                  "Salary Payment",
				ResultURL:                "invalid-url", // Invalid URL
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "SecurityCredential Too Short",
			req: AccountBalanceRequest{
				CommandID:                types.AccountBalanceCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				Initiator:                "apiuser",
				PartyA:                   600000,
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				Remarks:                  "Salary Payment",
				ResultURL:                "https://yourdomain.com/result",
				SecurityCredential:       "short", // Less than 8 characters
				OriginatorConversationID: "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Remarks Too Long",
			req: AccountBalanceRequest{
				CommandID:                types.AccountBalanceCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				Initiator:                "apiuser",
				PartyA:                   600000,
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				Remarks:                  string(make([]byte, 501)), // 501 characters
				ResultURL:                "https://yourdomain.com/result",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "validOriginatorConversationID",
			},
			wantErr: true,
		},
		{
			name: "Missing OriginatorConversationID",
			req: AccountBalanceRequest{
				CommandID:                types.AccountBalanceCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				Initiator:                "apiuser",
				PartyA:                   600000,
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				Remarks:                  "Salary Payment",
				ResultURL:                "https://yourdomain.com/result",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "", // Missing
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            err := tt.req.Validate(v)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

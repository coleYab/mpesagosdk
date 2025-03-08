package b2c

import (
	"testing"

	"github.com/coleYab/mpesasdk/types"
	"github.com/go-playground/validator/v10"
)

func TestB2CRequestValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		req     B2CRequest
		wantErr bool
	}{
		// 1. Valid Input
		{
			name: "Valid Input",
			req: B2CRequest{
				InitiatorName:            "apiuser",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "conv12345",
				CommandID:                types.BusinessPaymentCommand,
				Occasion:                 "Occasion",
				Amount:                   1030,
				PartyA:                   600000,
				PartyB:                   251700404709,
				Remarks:                  "Salary Payment",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				ResultURL:                "https://yourdomain.com/result",
			},
			wantErr: false,
		},

		// 2. Missing Required Fields
		{
			name: "Missing InitiatorName",
			req: B2CRequest{
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "conv12345",
				CommandID:                types.BusinessPaymentCommand,
				Occasion:                 "Occasion",
				Amount:                   1030,
				PartyA:                   600000,
				PartyB:                   251700404709,
				Remarks:                  "Salary Payment",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				ResultURL:                "https://yourdomain.com/result",
			},
			wantErr: true,
		},

		// 3. Invalid CommandID
		{
			name: "Invalid CommandID",
			req: B2CRequest{
				InitiatorName:            "apiuser",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "conv12345",
				CommandID:                types.CommandId(""), // Empty CommandID
				Occasion:                 "Occasion",
				Amount:                   1030,
				PartyA:                   600000,
				PartyB:                   251700404709,
				Remarks:                  "Salary Payment",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				ResultURL:                "https://yourdomain.com/result",
			},
			wantErr: true,
		},

		// 4. Invalid Amount
		{
			name: "Invalid Amount (Less than 1)",
			req: B2CRequest{
				InitiatorName:            "apiuser",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "conv12345",
				CommandID:                types.BusinessPaymentCommand,
				Occasion:                 "Occasion",
				Amount:                   0, // Invalid Amount
				PartyA:                   600000,
				PartyB:                   251700404709,
				Remarks:                  "Salary Payment",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				ResultURL:                "https://yourdomain.com/result",
			},
			wantErr: true,
		},

		// 5. Invalid PartyA
		{
			name: "Missing PartyA",
			req: B2CRequest{
				InitiatorName:            "apiuser",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "conv12345",
				CommandID:                types.BusinessPaymentCommand,
				Occasion:                 "Occasion",
				Amount:                   1030,
				PartyB:                   251700404709,
				Remarks:                  "Salary Payment",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				ResultURL:                "https://yourdomain.com/result",
			},
			wantErr: true,
		},

		// 6. Invalid PartyB
		{
			name: "Missing PartyB",
			req: B2CRequest{
				InitiatorName:            "apiuser",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "conv12345",
				CommandID:                types.BusinessPaymentCommand,
				Occasion:                 "Occasion",
				Amount:                   1030,
				PartyA:                   600000,
				Remarks:                  "Salary Payment",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				ResultURL:                "https://yourdomain.com/result",
			},
			wantErr: true,
		},

		// 7. Remarks Too Long
		{
			name: "Remarks Too Long",
			req: B2CRequest{
				InitiatorName:            "apiuser",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "conv12345",
				CommandID:                types.BusinessPaymentCommand,
				Occasion:                 "Occasion",
				Amount:                   1030,
				PartyA:                   600000,
				PartyB:                   251700404709,
				Remarks:                  string(make([]byte, 201)), // 201 characters
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				ResultURL:                "https://yourdomain.com/result",
			},
			wantErr: true,
		},

		// 8. Invalid QueueTimeOutURL
		{
			name: "Invalid QueueTimeOutURL",
			req: B2CRequest{
				InitiatorName:            "apiuser",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "conv12345",
				CommandID:                types.BusinessPaymentCommand,
				Occasion:                 "Occasion",
				Amount:                   1030,
				PartyA:                   600000,
				PartyB:                   251700404709,
				Remarks:                  "Salary Payment",
				QueueTimeOutURL:          "invalid-url", // Invalid URL
				ResultURL:                "https://yourdomain.com/result",
			},
			wantErr: true,
		},

		// 9. Invalid ResultURL
		{
			name: "Invalid ResultURL",
			req: B2CRequest{
				InitiatorName:            "apiuser",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				OriginatorConversationID: "conv12345",
				CommandID:                types.BusinessPaymentCommand,
				Occasion:                 "Occasion",
				Amount:                   1030,
				PartyA:                   600000,
				PartyB:                   251700404709,
				Remarks:                  "Salary Payment",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
				ResultURL:                "invalid-url", // Invalid URL
			},
			wantErr: true,
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate the B2CRequest
			err := tt.req.Validate(validate)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err != nil)
			}
		})
	}
}

package transaction

import (
	"testing"

	"github.com/coleYab/mpesagosdk/types"
	"github.com/go-playground/validator/v10"
)

func TestTransactionReversalRequestValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		req     TransactionReversalRequest
		wantErr bool
	}{
		// 1. Valid Input
		{
			name: "Valid Input",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: false,
		},

		// 2. Missing Initiator
		{
			name: "Missing Initiator",
			req: TransactionReversalRequest{
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},

		// 3. Missing TransactionID
		{
			name: "Missing TransactionID",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},

		// 4. Missing SecurityCredential
		{
			name: "Missing SecurityCredential",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},

		// 5. Missing PartyA
		{
			name: "Missing PartyA",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},

		// 6. Missing CommandID
		{
			name: "Missing CommandID",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},

		// 7. Missing IdentifierType
		{
			name: "Missing IdentifierType",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},

		// 8. Missing Amount
		{
			name: "Missing Amount",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},

		// 9. Invalid Amount (Less than 1)
		{
			name: "Invalid Amount (Less than 1)",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   0, // Invalid amount, should be greater than 0
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},

		// 10. Invalid ResultURL
		{
			name: "Invalid ResultURL",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  "Reversing transaction",
				ResultURL:                "invalid-url", // Invalid URL format
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},

		// 11. Invalid QueueTimeOutURL
		{
			name: "Invalid QueueTimeOutURL",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  "Reversing transaction",
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "invalid-url", // Invalid URL format
			},
			wantErr: true,
		},

		// 12. Remarks Too Long
		{
			name: "Remarks Too Long",
			req: TransactionReversalRequest{
				Initiator:                "appuser",
				TransactionID:            "LKXXXX1234",
				SecurityCredential:       "PU8f0AptZr16W28uzZy8+Ke4ww+HDk6/WXGurNcKREm7ihjUHL0TGWBxWbIzhftZkEms6LHhZlzh36LtAjLLxLiCRXHIW5Fv6oqOIsrl9pMw0F5pfEPMzDEXNlotjMpaFcEFS1GpnHWkIOaguXMNaf0Uev49rjzER495LMP3Z9EIPJmOuOI5QUZ6h3udctyyKIeUBdab0vf0zATY66Zm9XZc2CHHx3NsyU7i680s1OWreZ7SobuXsEyjZlh4hb1G0HNICFt/kp0PZN8Pt09qBeLX5BE1Tre0bb4v66AatJEuXQA39VJCZ6A+UldKyb5HLsdQHn+eZvd/K2yLtwpCxA==",
				PartyA:                   "600000",
				CommandID:                types.TransactionReversalCommand,
				IdentifierType:           types.ShortCodeIdentifierType,
				OriginatorConversationID: "conv12345",
				Amount:                   1000,
				Remarks:                  string(make([]byte, 201)), // 201 characters, exceeds limit
				ResultURL:                "https://yourdomain.com/result",
				QueueTimeOutURL:          "https://yourdomain.com/timeout",
			},
			wantErr: true,
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate the TransactionReversalRequest
			err := validate.Struct(tt.req)

			// Check if we expect an error
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v, message: %v", tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

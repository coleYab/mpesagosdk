package transaction

import (
	"testing"

	"github.com/coleYab/mpesagosdk/types"
	"github.com/go-playground/validator/v10"
)

func TestTransactionStatusRequestValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		req     TransactionStatusRequest
		wantErr bool
	}{
		// 1. Valid Input
		{
			name: "Valid Input",
			req: TransactionStatusRequest{
				Initiator:                "apitest",
				SecurityCredential:       "lMhf0UqE4ydeEDwpUskmPgkNDZnA6NLi7z3T1TQuWCkH3/ScW8pRRnobq/AcwFvbC961+zDMgOEYGm8Oivb7L/7Y9ED3lhR7pJvnH8B1wYis5ifdeeWI6XE2NSq8X1Tc7QB9Dg8SlPEud3tgloB2DlT+JIv3ebIl/J/8ihGVrq499bt1pz/EA2nzkCtGeHRNbEDxkqkEnbioV0OM//0bv4K++XyV6jUFlIIgkDkmcK6aOU8mPBHs2um9aP+Y+nTJaa6uHDudRFg0+3G6gt1zRCPs8AYbts2IebseBGfZKv5K6Lqk9/W8657gEkrDZE8Mi78MVianqHdY/8d6D9KKhw==",
				CommandID:                types.TransactionStatusCommand,
				TransactionID:            "0",
				OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
				PartyA:                   "1020",
				IdentifierType:           types.IdentifierType("4"),
				ResultURL:                "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				QueueTimeOutURL:          "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				Remarks:                  "Trans Status",
				Occasion:                 "Query trans status",
			},
			wantErr: false,
		},

		// 2. Missing Initiator
		{
			name: "Missing Initiator",
			req: TransactionStatusRequest{
				SecurityCredential:       "lMhf0UqE4ydeEDwpUskmPgkNDZnA6NLi7z3T1TQuWCkH3/ScW8pRRnobq/AcwFvbC961+zDMgOEYGm8Oivb7L/7Y9ED3lhR7pJvnH8B1wYis5ifdeeWI6XE2NSq8X1Tc7QB9Dg8SlPEud3tgloB2DlT+JIv3ebIl/J/8ihGVrq499bt1pz/EA2nzkCtGeHRNbEDxkqkEnbioV0OM//0bv4K++XyV6jUFlIIgkDkmcK6aOU8mPBHs2um9aP+Y+nTJaa6uHDudRFg0+3G6gt1zRCPs8AYbts2IebseBGfZKv5K6Lqk9/W8657gEkrDZE8Mi78MVianqHdY/8d6D9KKhw==",
				CommandID:                types.TransactionStatusCommand,
				TransactionID:            "0",
				OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
				PartyA:                   "1020",
				IdentifierType:           types.IdentifierType("4"),
				ResultURL:                "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				QueueTimeOutURL:          "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				Remarks:                  "Trans Status",
				Occasion:                 "Query trans status",
			},
			wantErr: true,
		},

		// 3. Missing SecurityCredential
		{
			name: "Missing SecurityCredential",
			req: TransactionStatusRequest{
				Initiator:                "apitest",
				CommandID:                types.TransactionStatusCommand,
				TransactionID:            "0",
				OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
				PartyA:                   "1020",
				IdentifierType:           types.IdentifierType("4"),
				ResultURL:                "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				QueueTimeOutURL:          "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				Remarks:                  "Trans Status",
				Occasion:                 "Query trans status",
			},
			wantErr: true,
		},

		// 4. Missing CommandID
		{
			name: "Missing CommandID",
			req: TransactionStatusRequest{
				Initiator:                "apitest",
				SecurityCredential:       "lMhf0UqE4ydeEDwpUskmPgkNDZnA6NLi7z3T1TQuWCkH3/ScW8pRRnobq/AcwFvbC961+zDMgOEYGm8Oivb7L/7Y9ED3lhR7pJvnH8B1wYis5ifdeeWI6XE2NSq8X1Tc7QB9Dg8SlPEud3tgloB2DlT+JIv3ebIl/J/8ihGVrq499bt1pz/EA2nzkCtGeHRNbEDxkqkEnbioV0OM//0bv4K++XyV6jUFlIIgkDkmcK6aOU8mPBHs2um9aP+Y+nTJaa6uHDudRFg0+3G6gt1zRCPs8AYbts2IebseBGfZKv5K6Lqk9/W8657gEkrDZE8Mi78MVianqHdY/8d6D9KKhw==",
				TransactionID:            "0",
				OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
				PartyA:                   "1020",
				IdentifierType:           types.IdentifierType("4"),
				ResultURL:                "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				QueueTimeOutURL:          "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				Remarks:                  "Trans Status",
				Occasion:                 "Query trans status",
			},
			wantErr: true,
		},

		// 5. Missing TransactionID
		{
			name: "Missing TransactionID",
			req: TransactionStatusRequest{
				Initiator:                "apitest",
				SecurityCredential:       "lMhf0UqE4ydeEDwpUskmPgkNDZnA6NLi7z3T1TQuWCkH3/ScW8pRRnobq/AcwFvbC961+zDMgOEYGm8Oivb7L/7Y9ED3lhR7pJvnH8B1wYis5ifdeeWI6XE2NSq8X1Tc7QB9Dg8SlPEud3tgloB2DlT+JIv3ebIl/J/8ihGVrq499bt1pz/EA2nzkCtGeHRNbEDxkqkEnbioV0OM//0bv4K++XyV6jUFlIIgkDkmcK6aOU8mPBHs2um9aP+Y+nTJaa6uHDudRFg0+3G6gt1zRCPs8AYbts2IebseBGfZKv5K6Lqk9/W8657gEkrDZE8Mi78MVianqHdY/8d6D9KKhw==",
				CommandID:                types.TransactionStatusCommand,
				OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
				PartyA:                   "1020",
				IdentifierType:           types.IdentifierType("4"),
				ResultURL:                "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				QueueTimeOutURL:          "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				Remarks:                  "Trans Status",
				Occasion:                 "Query trans status",
			},
			wantErr: true,
		},

		// 6. Missing PartyA
		{
			name: "Missing PartyA",
			req: TransactionStatusRequest{
				Initiator:                "apitest",
				SecurityCredential:       "lMhf0UqE4ydeEDwpUskmPgkNDZnA6NLi7z3T1TQuWCkH3/ScW8pRRnobq/AcwFvbC961+zDMgOEYGm8Oivb7L/7Y9ED3lhR7pJvnH8B1wYis5ifdeeWI6XE2NSq8X1Tc7QB9Dg8SlPEud3tgloB2DlT+JIv3ebIl/J/8ihGVrq499bt1pz/EA2nzkCtGeHRNbEDxkqkEnbioV0OM//0bv4K++XyV6jUFlIIgkDkmcK6aOU8mPBHs2um9aP+Y+nTJaa6uHDudRFg0+3G6gt1zRCPs8AYbts2IebseBGfZKv5K6Lqk9/W8657gEkrDZE8Mi78MVianqHdY/8d6D9KKhw==",
				CommandID:                types.TransactionStatusCommand,
				TransactionID:            "0",
				OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
				IdentifierType:           types.IdentifierType("4"),
				ResultURL:                "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				QueueTimeOutURL:          "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				Remarks:                  "Trans Status",
				Occasion:                 "Query trans status",
			},
			wantErr: true,
		},

		// 7. Invalid ResultURL
		{
			name: "Invalid ResultURL",
			req: TransactionStatusRequest{
				Initiator:                "apitest",
				SecurityCredential:       "lMhf0UqE4ydeEDwpUskmPgkNDZnA6NLi7z3T1TQuWCkH3/ScW8pRRnobq/AcwFvbC961+zDMgOEYGm8Oivb7L/7Y9ED3lhR7pJvnH8B1wYis5ifdeeWI6XE2NSq8X1Tc7QB9Dg8SlPEud3tgloB2DlT+JIv3ebIl/J/8ihGVrq499bt1pz/EA2nzkCtGeHRNbEDxkqkEnbioV0OM//0bv4K++XyV6jUFlIIgkDkmcK6aOU8mPBHs2um9aP+Y+nTJaa6uHDudRFg0+3G6gt1zRCPs8AYbts2IebseBGfZKv5K6Lqk9/W8657gEkrDZE8Mi78MVianqHdY/8d6D9KKhw==",
				CommandID:                types.TransactionStatusCommand,
				TransactionID:            "0",
				OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
				PartyA:                   "1020",
				IdentifierType:           types.IdentifierType("4"),
				ResultURL:                "invalid-url", // Invalid URL
				QueueTimeOutURL:          "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				Remarks:                  "Trans Status",
				Occasion:                 "Query trans status",
			},
			wantErr: true,
		},

		// 8. Invalid QueueTimeOutURL
		{
			name: "Invalid QueueTimeOutURL",
			req: TransactionStatusRequest{
				Initiator:                "apitest",
				SecurityCredential:       "lMhf0UqE4ydeEDwpUskmPgkNDZnA6NLi7z3T1TQuWCkH3/ScW8pRRnobq/AcwFvbC961+zDMgOEYGm8Oivb7L/7Y9ED3lhR7pJvnH8B1wYis5ifdeeWI6XE2NSq8X1Tc7QB9Dg8SlPEud3tgloB2DlT+JIv3ebIl/J/8ihGVrq499bt1pz/EA2nzkCtGeHRNbEDxkqkEnbioV0OM//0bv4K++XyV6jUFlIIgkDkmcK6aOU8mPBHs2um9aP+Y+nTJaa6uHDudRFg0+3G6gt1zRCPs8AYbts2IebseBGfZKv5K6Lqk9/W8657gEkrDZE8Mi78MVianqHdY/8d6D9KKhw==",
				CommandID:                types.TransactionStatusCommand,
				TransactionID:            "0",
				OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
				PartyA:                   "1020",
				IdentifierType:           types.IdentifierType("4"),
				ResultURL:                "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				QueueTimeOutURL:          "invalid-url", // Invalid URL
				Remarks:                  "Trans Status",
				Occasion:                 "Query trans status",
			},
			wantErr: true,
		},

		// 9. Remarks Too Long
		{
			name: "Remarks Too Long",
			req: TransactionStatusRequest{
				Initiator:                "apitest",
				SecurityCredential:       "lMhf0UqE4ydeEDwpUskmPgkNDZnA6NLi7z3T1TQuWCkH3/ScW8pRRnobq/AcwFvbC961+zDMgOEYGm8Oivb7L/7Y9ED3lhR7pJvnH8B1wYis5ifdeeWI6XE2NSq8X1Tc7QB9Dg8SlPEud3tgloB2DlT+JIv3ebIl/J/8ihGVrq499bt1pz/EA2nzkCtGeHRNbEDxkqkEnbioV0OM//0bv4K++XyV6jUFlIIgkDkmcK6aOU8mPBHs2um9aP+Y+nTJaa6uHDudRFg0+3G6gt1zRCPs8AYbts2IebseBGfZKv5K6Lqk9/W8657gEkrDZE8Mi78MVianqHdY/8d6D9KKhw==",
				CommandID:                types.TransactionStatusCommand,
				TransactionID:            "0",
				OriginatorConversationID: "AG-20190826-0000777ab7d848b9e721",
				PartyA:                   "1020",
				IdentifierType:           types.IdentifierType("4"),
				ResultURL:                "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				QueueTimeOutURL:          "https://webhook.site/7ed4b055-fa4d-45f3-ae1f-328c52aa4d7d",
				Remarks:                  string(make([]byte, 501)), // 501 characters, exceeds 500 character limit
				Occasion:                 "Query trans status",
			},
			wantErr: true,
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate the TransactionStatusRequest
			err := validate.Struct(tt.req)

			// Check if we expect an error
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err != nil)
			}
		})
	}
}

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
				InitiatorName:            "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				Amount:                   100,
				PartyA:                   123456,
				PartyB:                   654321,
				Remarks:                  "Payment for services",
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Occasion:                 "Birthday",
				OriginatorConversationID: "conv12345",
			},
			wantErr: false,
		},

		// 2. Missing Required Fields
		{
			name: "Missing InitiatorName",
			req: B2CRequest{
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				Amount:                   100,
				PartyA:                   123456,
				PartyB:                   654321,
				Remarks:                  "Payment for services",
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Occasion:                 "Birthday",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 3. Invalid CommandID
		{
			name: "Invalid CommandID",
			req: B2CRequest{
				InitiatorName:            "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId(""), // Empty CommandID
				Amount:                   100,
				PartyA:                   123456,
				PartyB:                   654321,
				Remarks:                  "Payment for services",
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Occasion:                 "Birthday",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 4. Invalid Amount
		{
			name: "Invalid Amount (Less than 1)",
			req: B2CRequest{
				InitiatorName:            "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				Amount:                   0, // Invalid Amount
				PartyA:                   123456,
				PartyB:                   654321,
				Remarks:                  "Payment for services",
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Occasion:                 "Birthday",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 5. Invalid PartyA
		{
			name: "Missing PartyA",
			req: B2CRequest{
				InitiatorName:            "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				Amount:                   100,
				PartyB:                   654321,
				Remarks:                  "Payment for services",
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Occasion:                 "Birthday",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 6. Invalid PartyB
		{
			name: "Missing PartyB",
			req: B2CRequest{
				InitiatorName:            "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				Amount:                   100,
				PartyA:                   123456,
				Remarks:                  "Payment for services",
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Occasion:                 "Birthday",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 7. Remarks Too Long
		{
			name: "Remarks Too Long",
			req: B2CRequest{
				InitiatorName:            "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				Amount:                   100,
				PartyA:                   123456,
				PartyB:                   654321,
				Remarks:                  string(make([]byte, 201)), // 201 characters
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Occasion:                 "Birthday",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 8. Invalid QueueTimeOutURL
		{
			name: "Invalid QueueTimeOutURL",
			req: B2CRequest{
				InitiatorName:            "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				Amount:                   100,
				PartyA:                   123456,
				PartyB:                   654321,
				Remarks:                  "Payment for services",
				QueueTimeOutURL:          "invalid-url", // Invalid URL
				ResultURL:                "https://example.com/result",
				Occasion:                 "Birthday",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 9. Invalid ResultURL
		{
			name: "Invalid ResultURL",
			req: B2CRequest{
				InitiatorName:            "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				Amount:                   100,
				PartyA:                   123456,
				PartyB:                   654321,
				Remarks:                  "Payment for services",
				QueueTimeOutURL:          "invalid-url", // Invalid URL
				ResultURL:                "https://example.com/result",
				Occasion:                 "Birthday",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 10. Invalid Occasion
		{
			name: "Invalid Occasion",
			req: B2CRequest{
				InitiatorName:            "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				Amount:                   100,
				PartyA:                   123456,
				PartyB:                   654321,
				Remarks:                  "Payment for services",
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Occasion:                 string(make([]byte, 101)), // Occasion too long, 101 characters
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate the B2CRequest
			err := validate.Struct(tt.req)

			// Check if we expect an error
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err != nil)
			}
		})
	}
}

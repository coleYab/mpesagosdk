package transaction

import (
	"testing"

	"github.com/coleYab/mpesasdk/types"
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
				Initiator:                "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				TransactionID:            "validTransactionID123",
				Amount:                   100,
				PartyA:                   "ReceiverParty",
				IdentifierType:           types.IdentifierType("MSISDN"), // Adjust this type based on your enum values
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "Reversal for payment",
				Occasion:                 "Refund",
				OriginatorConversationID: "conv12345",
			},
			wantErr: false,
		},

		// 2. Missing Initiator
		{
			name: "Missing Initiator",
			req: TransactionReversalRequest{
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				TransactionID:            "validTransactionID123",
				Amount:                   100,
				PartyA:                   "ReceiverParty",
				IdentifierType:           types.IdentifierType("MSISDN"),
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "Reversal for payment",
				Occasion:                 "Refund",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 3. Invalid TransactionID
		{
			name: "Invalid TransactionID (Too short)",
			req: TransactionReversalRequest{
				Initiator:                "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				TransactionID:            "short", // Too short, should be at least 10 characters
				Amount:                   100,
				PartyA:                   "ReceiverParty",
				IdentifierType:           types.IdentifierType("MSISDN"),
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "Reversal for payment",
				Occasion:                 "Refund",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 4. Invalid Amount (Less than 1)
		{
			name: "Invalid Amount (Less than 1)",
			req: TransactionReversalRequest{
				Initiator:                "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				TransactionID:            "validTransactionID123",
				Amount:                   0, // Amount must be greater than 0
				PartyA:                   "ReceiverParty",
				IdentifierType:           types.IdentifierType("MSISDN"),
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "Reversal for payment",
				Occasion:                 "Refund",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 5. Missing PartyA
		{
			name: "Missing PartyA",
			req: TransactionReversalRequest{
				Initiator:                "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				TransactionID:            "validTransactionID123",
				Amount:                   100,
				IdentifierType:           types.IdentifierType("MSISDN"),
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "Reversal for payment",
				Occasion:                 "Refund",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 6. Invalid QueueTimeOutURL
		{
			name: "Invalid QueueTimeOutURL",
			req: TransactionReversalRequest{
				Initiator:                "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				TransactionID:            "validTransactionID123",
				Amount:                   100,
				PartyA:                   "ReceiverParty",
				IdentifierType:           types.IdentifierType("MSISDN"),
				QueueTimeOutURL:          "invalid-url", // Invalid URL format
				ResultURL:                "https://example.com/result",
				Remarks:                  "Reversal for payment",
				Occasion:                 "Refund",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 7. Invalid ResultURL
		{
			name: "Invalid ResultURL",
			req: TransactionReversalRequest{
				Initiator:                "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				TransactionID:            "validTransactionID123",
				Amount:                   100,
				PartyA:                   "ReceiverParty",
				IdentifierType:           types.IdentifierType("MSISDN"),
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "invalid-url", // Invalid URL format
				Remarks:                  "Reversal for payment",
				Occasion:                 "Refund",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 8. Remarks Too Long
		{
			name: "Remarks Too Long",
			req: TransactionReversalRequest{
				Initiator:                "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				TransactionID:            "validTransactionID123",
				Amount:                   100,
				PartyA:                   "ReceiverParty",
				IdentifierType:           types.IdentifierType("MSISDN"),
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  string(make([]byte, 201)), // 201 characters, exceeds limit
				Occasion:                 "Refund",
				OriginatorConversationID: "conv12345",
			},
			wantErr: true,
		},

		// 9. Invalid Occasion
		{
			name: "Invalid Occasion",
			req: TransactionReversalRequest{
				Initiator:                "JohnDoe",
				SecurityCredential:       "secureCredential123",
				CommandID:                types.CommandId("ValidCommandID"),
				TransactionID:            "validTransactionID123",
				Amount:                   100,
				PartyA:                   "ReceiverParty",
				IdentifierType:           types.IdentifierType("MSISDN"),
				QueueTimeOutURL:          "https://example.com/timeout",
				ResultURL:                "https://example.com/result",
				Remarks:                  "Reversal for payment",
				Occasion:                 string(make([]byte, 101)), // Occasion too long (max length is 100)
				OriginatorConversationID: "conv12345",
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
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err != nil)
			}
		})
	}
}

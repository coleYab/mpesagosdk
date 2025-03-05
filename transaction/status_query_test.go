package transaction

import (
	"testing"

	"github.com/coleYab/mpesasdk/types"
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
				CommandID:                types.CommandId("ValidCommandID"),
				IdentifierType:           types.IdentifierType("MSISDN"), // Adjust this type based on your enum values
				Initiator:                "JohnDoe",
				Occasion:                 "Payment inquiry",
				OriginatorConversationID: "conv12345",
				PartyA:                   "ReceiverParty",
				QueueTimeOutURL:          "https://example.com/timeout",
				Remarks:                  "Transaction status check",
				ResultURL:                "https://example.com/result",
				SecurityCredential:       "secureCredential123",
				TransactionID:            "validTransactionID123",
			},
			wantErr: false,
		},

		// 2. Missing CommandID
		{
			name: "Missing CommandID",
			req: TransactionStatusRequest{
				IdentifierType:           types.IdentifierType("MSISDN"),
				Initiator:                "JohnDoe",
				Occasion:                 "Payment inquiry",
				OriginatorConversationID: "conv12345",
				PartyA:                   "ReceiverParty",
				QueueTimeOutURL:          "https://example.com/timeout",
				Remarks:                  "Transaction status check",
				ResultURL:                "https://example.com/result",
				SecurityCredential:       "secureCredential123",
				TransactionID:            "validTransactionID123",
			},
			wantErr: true,
		},

		// 3. Missing IdentifierType
		{
			name: "Missing IdentifierType",
			req: TransactionStatusRequest{
				CommandID:                types.CommandId("ValidCommandID"),
				Initiator:                "JohnDoe",
				Occasion:                 "Payment inquiry",
				OriginatorConversationID: "conv12345",
				PartyA:                   "ReceiverParty",
				QueueTimeOutURL:          "https://example.com/timeout",
				Remarks:                  "Transaction status check",
				ResultURL:                "https://example.com/result",
				SecurityCredential:       "secureCredential123",
				TransactionID:            "validTransactionID123",
			},
			wantErr: true,
		},

		// 4. Missing PartyA
		{
			name: "Missing PartyA",
			req: TransactionStatusRequest{
				CommandID:                types.CommandId("ValidCommandID"),
				IdentifierType:           types.IdentifierType("MSISDN"),
				Initiator:                "JohnDoe",
				Occasion:                 "Payment inquiry",
				OriginatorConversationID: "conv12345",
				QueueTimeOutURL:          "https://example.com/timeout",
				Remarks:                  "Transaction status check",
				ResultURL:                "https://example.com/result",
				SecurityCredential:       "secureCredential123",
				TransactionID:            "validTransactionID123",
			},
			wantErr: true,
		},

		// 5. Invalid QueueTimeOutURL
		{
			name: "Invalid QueueTimeOutURL",
			req: TransactionStatusRequest{
				CommandID:                types.CommandId("ValidCommandID"),
				IdentifierType:           types.IdentifierType("MSISDN"),
				Initiator:                "JohnDoe",
				Occasion:                 "Payment inquiry",
				OriginatorConversationID: "conv12345",
				PartyA:                   "ReceiverParty",
				QueueTimeOutURL:          "invalid-url", // Invalid URL
				Remarks:                  "Transaction status check",
				ResultURL:                "https://example.com/result",
				SecurityCredential:       "secureCredential123",
				TransactionID:            "validTransactionID123",
			},
			wantErr: true,
		},

		// 6. Invalid ResultURL
		{
			name: "Invalid ResultURL",
			req: TransactionStatusRequest{
				CommandID:                types.CommandId("ValidCommandID"),
				IdentifierType:           types.IdentifierType("MSISDN"),
				Initiator:                "JohnDoe",
				Occasion:                 "Payment inquiry",
				OriginatorConversationID: "conv12345",
				PartyA:                   "ReceiverParty",
				QueueTimeOutURL:          "https://example.com/timeout",
				Remarks:                  "Transaction status check",
				ResultURL:                "invalid-url", // Invalid URL
				SecurityCredential:       "secureCredential123",
				TransactionID:            "validTransactionID123",
			},
			wantErr: true,
		},

		// 7. Missing SecurityCredential
		{
			name: "Missing SecurityCredential",
			req: TransactionStatusRequest{
				CommandID:                types.CommandId("ValidCommandID"),
				IdentifierType:           types.IdentifierType("MSISDN"),
				Initiator:                "JohnDoe",
				Occasion:                 "Payment inquiry",
				OriginatorConversationID: "conv12345",
				PartyA:                   "ReceiverParty",
				QueueTimeOutURL:          "https://example.com/timeout",
				Remarks:                  "Transaction status check",
				ResultURL:                "https://example.com/result",
				TransactionID:            "validTransactionID123",
			},
			wantErr: true,
		},

		// 8. Invalid TransactionID (Too Short)
		{
			name: "Invalid TransactionID (Too Short)",
			req: TransactionStatusRequest{
				CommandID:                types.CommandId("ValidCommandID"),
				IdentifierType:           types.IdentifierType("MSISDN"),
				Initiator:                "JohnDoe",
				Occasion:                 "Payment inquiry",
				OriginatorConversationID: "conv12345",
				PartyA:                   "ReceiverParty",
				QueueTimeOutURL:          "https://example.com/timeout",
				Remarks:                  "Transaction status check",
				ResultURL:                "https://example.com/result",
				SecurityCredential:       "secureCredential123",
				TransactionID:            "short", // Too short, should be at least 1 character
			},
			wantErr: true,
		},

		// 9. Remarks Too Long
		{
			name: "Remarks Too Long",
			req: TransactionStatusRequest{
				CommandID:                types.CommandId("ValidCommandID"),
				IdentifierType:           types.IdentifierType("MSISDN"),
				Initiator:                "JohnDoe",
				Occasion:                 "Payment inquiry",
				OriginatorConversationID: "conv12345",
				PartyA:                   "ReceiverParty",
				QueueTimeOutURL:          "https://example.com/timeout",
				Remarks:                  string(make([]byte, 501)), // 501 characters, exceeds 500 character limit
				ResultURL:                "https://example.com/result",
				SecurityCredential:       "secureCredential123",
				TransactionID:            "validTransactionID123",
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

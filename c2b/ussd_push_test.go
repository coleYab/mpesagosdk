package c2b

import (
	"testing"

	"github.com/coleYab/mpesasdk/types"
	"github.com/go-playground/validator/v10"
)

func TestUSSDPaymentRequestValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		req     USSDPaymentRequest
		wantErr bool
	}{
		// 1. Valid Input
		{
			name: "Valid Input",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "12345",
				ReferenceData:     []ReferenceDataRequest{{Key: "CustomerId", Value: "123"}},
				TransactionType:   types.TransactionType("Payment"),
				Password:          "securePassword123",
				Timestamp:         "20250305120000",
				Amount:            100,
				PartyA:            "PartyA123",
				PartyB:            "PartyB123",
				PhoneNumber:       "254701234567",
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "AccountRef123",
				TransactionDesc:   "Payment for goods",
			},
			wantErr: false,
		},

		// 2. Missing MerchantRequestID
		{
			name: "Missing MerchantRequestID",
			req: USSDPaymentRequest{
				BusinessShortCode: "12345",
				ReferenceData:     []ReferenceDataRequest{{Key: "CustomerId", Value: "123"}},
				TransactionType:   types.TransactionType("Payment"),
				Password:          "securePassword123",
				Timestamp:         "20250305120000",
				Amount:            100,
				PartyA:            "PartyA123",
				PartyB:            "PartyB123",
				PhoneNumber:       "254701234567",
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "AccountRef123",
				TransactionDesc:   "Payment for goods",
			},
			wantErr: true,
		},

		// 3. Invalid MerchantRequestID (Too short)
		{
			name: "Invalid MerchantRequestID (Too short)",
			req: USSDPaymentRequest{
				MerchantRequestID: "A", // Less than 1 character
				BusinessShortCode: "12345",
				ReferenceData:     []ReferenceDataRequest{{Key: "CustomerId", Value: "123"}},
				TransactionType:   types.TransactionType("Payment"),
				Password:          "securePassword123",
				Timestamp:         "20250305120000",
				Amount:            100,
				PartyA:            "PartyA123",
				PartyB:            "PartyB123",
				PhoneNumber:       "254701234567",
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "AccountRef123",
				TransactionDesc:   "Payment for goods",
			},
			wantErr: true,
		},

		// 4. Missing Amount
		{
			name: "Missing Amount",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "12345",
				ReferenceData:     []ReferenceDataRequest{{Key: "CustomerId", Value: "123"}},
				TransactionType:   types.TransactionType("Payment"),
				Password:          "securePassword123",
				Timestamp:         "20250305120000",
				PartyA:            "PartyA123",
				PartyB:            "PartyB123",
				PhoneNumber:       "254701234567",
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "AccountRef123",
				TransactionDesc:   "Payment for goods",
			},
			wantErr: true,
		},

		// 5. Invalid Amount (Less than 1)
		{
			name: "Invalid Amount (Less than 1)",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "12345",
				ReferenceData:     []ReferenceDataRequest{{Key: "CustomerId", Value: "123"}},
				TransactionType:   types.TransactionType("Payment"),
				Password:          "securePassword123",
				Timestamp:         "20250305120000",
				Amount:            0, // Invalid amount, should be greater than 0
				PartyA:            "PartyA123",
				PartyB:            "PartyB123",
				PhoneNumber:       "254701234567",
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "AccountRef123",
				TransactionDesc:   "Payment for goods",
			},
			wantErr: true,
		},

		// 6. Invalid PhoneNumber (Missing)
		{
			name: "Invalid PhoneNumber (Missing)",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "12345",
				ReferenceData:     []ReferenceDataRequest{{Key: "CustomerId", Value: "123"}},
				TransactionType:   types.TransactionType("Payment"),
				Password:          "securePassword123",
				Timestamp:         "20250305120000",
				Amount:            100,
				PartyA:            "PartyA123",
				PartyB:            "PartyB123",
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "AccountRef123",
				TransactionDesc:   "Payment for goods",
			},
			wantErr: true,
		},

		// 7. Invalid PhoneNumber (Non-numeric)
		{
			name: "Invalid PhoneNumber (Non-numeric)",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "12345",
				ReferenceData:     []ReferenceDataRequest{{Key: "CustomerId", Value: "123"}},
				TransactionType:   types.TransactionType("Payment"),
				Password:          "securePassword123",
				Timestamp:         "20250305120000",
				Amount:            100,
				PartyA:            "PartyA123",
				PartyB:            "PartyB123",
				PhoneNumber:       "25470A12345", // Invalid phone number
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "AccountRef123",
				TransactionDesc:   "Payment for goods",
			},
			wantErr: true,
		},

		// 8. Invalid Timestamp (Incorrect format)
		{
			name: "Invalid Timestamp (Incorrect format)",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "12345",
				ReferenceData:     []ReferenceDataRequest{{Key: "CustomerId", Value: "123"}},
				TransactionType:   types.TransactionType("Payment"),
				Password:          "securePassword123",
				Timestamp:         "2025-03-05 12:00:00", // Incorrect timestamp format
				Amount:            100,
				PartyA:            "PartyA123",
				PartyB:            "PartyB123",
				PhoneNumber:       "254701234567",
				CallBackURL:       "https://example.com/callback",
				AccountReference:  "AccountRef123",
				TransactionDesc:   "Payment for goods",
			},
			wantErr: true,
		},

		// 9. Invalid CallBackURL
		{
			name: "Invalid CallBackURL",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "12345",
				ReferenceData:     []ReferenceDataRequest{{Key: "CustomerId", Value: "123"}},
				TransactionType:   types.TransactionType("Payment"),
				Password:          "securePassword123",
				Timestamp:         "20250305120000",
				Amount:            100,
				PartyA:            "PartyA123",
				PartyB:            "PartyB123",
				PhoneNumber:       "254701234567",
				CallBackURL:       "invalid-url", // Invalid URL format
				AccountReference:  "AccountRef123",
				TransactionDesc:   "Payment for goods",
			},
			wantErr: true,
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate the USSDPaymentRequest struct
			err := validate.Struct(tt.req)

			// Check if we expect an error
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err != nil)
			}
		})
	}
}

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
				BusinessShortCode: "1020",
				Password:          "M2VkZGU2YWY1Y2RhMzIyOWRjMmFkMTRiMjdjOWIwOWUxZDFlZDZiNGQ0OGYyMDRiNjg0ZDZhNWM2NTQyNTk2ZA==",
				Timestamp:         "20240918055823",
				TransactionType:   types.TransactionType("CustomerPayBillOnline"),
				Amount:            20,
				PartyA:            "251710404709",
				PartyB:            "1020",
				PhoneNumber:       "251700404709",
				CallBackURL:       "https://www.myservice:8080/result",
				AccountReference:  "Partner Unique ID",
				TransactionDesc:   "Payment Reason",
				ReferenceData: []ReferenceDataRequest{
					{
						Key:   "ThirdPartyReference",
						Value: "Ref-12345",
					},
				},
			},
			wantErr: false,
		},

		// 2. Missing MerchantRequestID
		{
			name: "Missing MerchantRequestID",
			req: USSDPaymentRequest{
				BusinessShortCode: "1020",
				Password:          "M2VkZGU2YWY1Y2RhMzIyOWRjMmFkMTRiMjdjOWIwOWUxZDFlZDZiNGQ0OGYyMDRiNjg0ZDZhNWM2NTQyNTk2ZA==",
				Timestamp:         "20240918055823",
				TransactionType:   types.TransactionType("CustomerPayBillOnline"),
				Amount:            20,
				PartyA:            "251710404709",
				PartyB:            "1020",
				PhoneNumber:       "251700404709",
				CallBackURL:       "https://www.myservice:8080/result",
				AccountReference:  "Partner Unique ID",
				TransactionDesc:   "Payment Reason",
				ReferenceData: []ReferenceDataRequest{
					{
						Key:   "ThirdPartyReference",
						Value: "Ref-12345",
					},
				},
			},
			wantErr: true,
		},

		{
			name: "Missing Amount",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "1020",
				Password:          "M2VkZGU2YWY1Y2RhMzIyOWRjMmFkMTRiMjdjOWIwOWUxZDFlZDZiNGQ0OGYyMDRiNjg0ZDZhNWM2NTQyNTk2ZA==",
				Timestamp:         "20240918055823",
				TransactionType:   types.TransactionType("CustomerPayBillOnline"),
				PartyA:            "251710404709",
				PartyB:            "1020",
				PhoneNumber:       "251700404709",
				CallBackURL:       "https://www.myservice:8080/result",
				AccountReference:  "Partner Unique ID",
				TransactionDesc:   "Payment Reason",
				ReferenceData: []ReferenceDataRequest{
					{
						Key:   "ThirdPartyReference",
						Value: "Ref-12345",
					},
				},
			},
			wantErr: true,
		},

		// 5. Invalid Amount (Less than 1)
		{
			name: "Invalid Amount (Less than 1)",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "1020",
				Password:          "M2VkZGU2YWY1Y2RhMzIyOWRjMmFkMTRiMjdjOWIwOWUxZDFlZDZiNGQ0OGYyMDRiNjg0ZDZhNWM2NTQyNTk2ZA==",
				Timestamp:         "20240918055823",
				TransactionType:   types.TransactionType("CustomerPayBillOnline"),
				Amount:            0, // Invalid amount, should be greater than 0
				PartyA:            "251710404709",
				PartyB:            "1020",
				PhoneNumber:       "251700404709",
				CallBackURL:       "https://www.myservice:8080/result",
				AccountReference:  "Partner Unique ID",
				TransactionDesc:   "Payment Reason",
				ReferenceData: []ReferenceDataRequest{
					{
						Key:   "ThirdPartyReference",
						Value: "Ref-12345",
					},
				},
			},
			wantErr: true,
		},

		// 6. Invalid PhoneNumber (Missing)
		{
			name: "Invalid PhoneNumber (Missing)",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "1020",
				Password:          "M2VkZGU2YWY1Y2RhMzIyOWRjMmFkMTRiMjdjOWIwOWUxZDFlZDZiNGQ0OGYyMDRiNjg0ZDZhNWM2NTQyNTk2ZA==",
				Timestamp:         "20240918055823",
				TransactionType:   types.TransactionType("CustomerPayBillOnline"),
				Amount:            20,
				PartyA:            "251710404709",
				PartyB:            "1020",
				CallBackURL:       "https://www.myservice:8080/result",
				AccountReference:  "Partner Unique ID",
				TransactionDesc:   "Payment Reason",
				ReferenceData: []ReferenceDataRequest{
					{
						Key:   "ThirdPartyReference",
						Value: "Ref-12345",
					},
				},
			},
			wantErr: true,
		},

		// 7. Invalid PhoneNumber (Non-numeric)
		{
			name: "Invalid PhoneNumber (Non-numeric)",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "1020",
				Password:          "M2VkZGU2YWY1Y2RhMzIyOWRjMmFkMTRiMjdjOWIwOWUxZDFlZDZiNGQ0OGYyMDRiNjg0ZDZhNWM2NTQyNTk2ZA==",
				Timestamp:         "20240918055823",
				TransactionType:   types.TransactionType("CustomerPayBillOnline"),
				Amount:            20,
				PartyA:            "251710404709",
				PartyB:            "1020",
				PhoneNumber:       "25170A12345", // Invalid phone number
				CallBackURL:       "https://www.myservice:8080/result",
				AccountReference:  "Partner Unique ID",
				TransactionDesc:   "Payment Reason",
				ReferenceData: []ReferenceDataRequest{
					{
						Key:   "ThirdPartyReference",
						Value: "Ref-12345",
					},
				},
			},
			wantErr: true,
		},

		// 8. Invalid Timestamp (Incorrect format)
		{
			name: "Invalid Timestamp (Incorrect format)",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "1020",
				Password:          "M2VkZGU2YWY1Y2RhMzIyOWRjMmFkMTRiMjdjOWIwOWUxZDFlZDZiNGQ0OGYyMDRiNjg0ZDZhNWM2NTQyNTk2ZA==",
				Timestamp:         "2024-03-05 12:00:00", // Incorrect timestamp format
				TransactionType:   types.TransactionType("CustomerPayBillOnline"),
				Amount:            20,
				PartyA:            "251710404709",
				PartyB:            "1020",
				PhoneNumber:       "251700404709",
				CallBackURL:       "https://www.myservice:8080/result",
				AccountReference:  "Partner Unique ID",
				TransactionDesc:   "Payment Reason",
				ReferenceData: []ReferenceDataRequest{
					{
						Key:   "ThirdPartyReference",
						Value: "Ref-12345",
					},
				},
			},
			wantErr: true,
		},

		// 9. Invalid CallBackURL
		{
			name: "Invalid CallBackURL",
			req: USSDPaymentRequest{
				MerchantRequestID: "MR12345",
				BusinessShortCode: "1020",
				Password:          "M2VkZGU2YWY1Y2RhMzIyOWRjMmFkMTRiMjdjOWIwOWUxZDFlZDZiNGQ0OGYyMDRiNjg0ZDZhNWM2NTQyNTk2ZA==",
				Timestamp:         "20240918055823",
				TransactionType:   types.TransactionType("CustomerPayBillOnline"),
				Amount:            20,
				PartyA:            "251710404709",
				PartyB:            "1020",
				PhoneNumber:       "251700404709",
				CallBackURL:       "invalid-url", // Invalid URL format
				AccountReference:  "Partner Unique ID",
				TransactionDesc:   "Payment Reason",
				ReferenceData: []ReferenceDataRequest{
					{
						Key:   "ThirdPartyReference",
						Value: "Ref-12345",
					},
				},
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

package c2b

import (
	"testing"

	"github.com/coleYab/mpesagosdk/types"
	"github.com/go-playground/validator/v10"
)

func TestSimulateCustomerInititatedPaymentValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		req     SimulateCustomerInititatedPayment
		wantErr bool
	}{
		// 1. Valid Input
		{
			name: "Valid Input",
			req: SimulateCustomerInititatedPayment{
				CommandID:     types.CommandId("SimulatePayment"),
				Amount:        100,
				Msisdn:        "254701234567", // Valid phone number format
				BillRefNumber: "123456789012", // Valid bill reference
				ShortCode:     "12345",
			},
			wantErr: false,
		},

		// 2. Missing CommandID
		{
			name: "Missing CommandID",
			req: SimulateCustomerInititatedPayment{
				Amount:        100,
				Msisdn:        "254701234567",
				BillRefNumber: "123456789012",
				ShortCode:     "12345",
			},
			wantErr: true,
		},

		// 3. Missing Amount
		{
			name: "Missing Amount",
			req: SimulateCustomerInititatedPayment{
				CommandID:     types.CommandId("SimulatePayment"),
				Msisdn:        "254701234567",
				BillRefNumber: "123456789012",
				ShortCode:     "12345",
			},
			wantErr: true,
		},

		// 4. Amount less than 1
		{
			name: "Amount Less Than 1",
			req: SimulateCustomerInititatedPayment{
				CommandID:     types.CommandId("SimulatePayment"),
				Amount:        0, // Invalid amount, should be greater than or equal to 1
				Msisdn:        "254701234567",
				BillRefNumber: "123456789012",
				ShortCode:     "12345",
			},
			wantErr: true,
		},

		// 5. Missing Msisdn
		{
			name: "Missing Msisdn",
			req: SimulateCustomerInititatedPayment{
				CommandID:     types.CommandId("SimulatePayment"),
				Amount:        100,
				BillRefNumber: "123456789012",
				ShortCode:     "12345",
			},
			wantErr: true,
		},

		// 6. Invalid Msisdn (not numeric)
		{
			name: "Invalid Msisdn",
			req: SimulateCustomerInititatedPayment{
				CommandID:     types.CommandId("SimulatePayment"),
				Amount:        100,
				Msisdn:        "25470A12345", // Invalid phone number (contains letters)
				BillRefNumber: "123456789012",
				ShortCode:     "12345",
			},
			wantErr: true,
		},

		// 7. Missing BillRefNumber
		{
			name: "Missing BillRefNumber",
			req: SimulateCustomerInititatedPayment{
				CommandID: types.CommandId("SimulatePayment"),
				Amount:    100,
				Msisdn:    "254701234567",
				ShortCode: "12345",
			},
			wantErr: true,
		},

		// 8. BillRefNumber Too Short
		{
			name: "BillRefNumber Too Short",
			req: SimulateCustomerInititatedPayment{
				CommandID:     types.CommandId("SimulatePayment"),
				Amount:        100,
				Msisdn:        "254701234567",
				BillRefNumber: "123", // Invalid bill reference number, should be at least 6 characters
				ShortCode:     "12345",
			},
			wantErr: true,
		},

		// 9. BillRefNumber Too Long
		{
			name: "BillRefNumber Too Long",
			req: SimulateCustomerInititatedPayment{
				CommandID:     types.CommandId("SimulatePayment"),
				Amount:        100,
				Msisdn:        "254701234567",
				BillRefNumber: "123456789012345678901", // Invalid bill reference number, should be at most 20 characters
				ShortCode:     "12345",
			},
			wantErr: true,
		},

		// 10. Missing ShortCode
		{
			name: "Missing ShortCode",
			req: SimulateCustomerInititatedPayment{
				CommandID:     types.CommandId("SimulatePayment"),
				Amount:        100,
				Msisdn:        "254701234567",
				BillRefNumber: "123456789012",
			},
			wantErr: true,
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate the SimulateCustomerInititatedPayment
			err := validate.Struct(tt.req)

			// Check if we expect an error
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err != nil)
			}
		})
	}
}

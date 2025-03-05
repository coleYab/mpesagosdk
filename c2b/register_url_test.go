package c2b

import (
	"testing"

	"github.com/coleYab/mpesasdk/types"
	"github.com/go-playground/validator/v10"
)

func TestRegisterC2BURLRequestValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		req     RegisterC2BURLRequest
		wantErr bool
	}{
		// 1. Valid Input
		{
			name: "Valid Input",
			req: RegisterC2BURLRequest{
				ShortCode:       "12345",
				ResponseType:    types.ResponseType("SUCCESS"),
				CommandID:       types.CommandId("RegisterC2B"),
				ConfirmationURL: "https://example.com/confirmation",
				ValidationURL:   "https://example.com/validation",
			},
			wantErr: false,
		},

		// 2. Missing ShortCode
		{
			name: "Missing ShortCode",
			req: RegisterC2BURLRequest{
				ResponseType:    types.ResponseType("SUCCESS"),
				CommandID:       types.CommandId("RegisterC2B"),
				ConfirmationURL: "https://example.com/confirmation",
				ValidationURL:   "https://example.com/validation",
			},
			wantErr: true,
		},

		// 3. ShortCode Too Short
		{
			name: "ShortCode Too Short",
			req: RegisterC2BURLRequest{
				ShortCode:       "1234", // Too short, should be at least 5 characters
				ResponseType:    types.ResponseType("SUCCESS"),
				CommandID:       types.CommandId("RegisterC2B"),
				ConfirmationURL: "https://example.com/confirmation",
				ValidationURL:   "https://example.com/validation",
			},
			wantErr: true,
		},

		// 4. ShortCode Too Long
		{
			name: "ShortCode Too Long",
			req: RegisterC2BURLRequest{
				ShortCode:       "123456789012345678901", // Too long, should be at most 20 characters
				ResponseType:    types.ResponseType("SUCCESS"),
				CommandID:       types.CommandId("RegisterC2B"),
				ConfirmationURL: "https://example.com/confirmation",
				ValidationURL:   "https://example.com/validation",
			},
			wantErr: true,
		},

		// 5. Missing ResponseType
		{
			name: "Missing ResponseType",
			req: RegisterC2BURLRequest{
				ShortCode:       "12345",
				CommandID:       types.CommandId("RegisterC2B"),
				ConfirmationURL: "https://example.com/confirmation",
				ValidationURL:   "https://example.com/validation",
			},
			wantErr: true,
		},

		// 6. Missing CommandID
		{
			name: "Missing CommandID",
			req: RegisterC2BURLRequest{
				ShortCode:       "12345",
				ResponseType:    types.ResponseType("SUCCESS"),
				ConfirmationURL: "https://example.com/confirmation",
				ValidationURL:   "https://example.com/validation",
			},
			wantErr: true,
		},

		// 7. Invalid ConfirmationURL
		{
			name: "Invalid ConfirmationURL",
			req: RegisterC2BURLRequest{
				ShortCode:       "12345",
				ResponseType:    types.ResponseType("SUCCESS"),
				CommandID:       types.CommandId("RegisterC2B"),
				ConfirmationURL: "invalid-url", // Invalid URL
				ValidationURL:   "https://example.com/validation",
			},
			wantErr: true,
		},

		// 8. Invalid ValidationURL
		{
			name: "Invalid ValidationURL",
			req: RegisterC2BURLRequest{
				ShortCode:       "12345",
				ResponseType:    types.ResponseType("SUCCESS"),
				CommandID:       types.CommandId("RegisterC2B"),
				ConfirmationURL: "https://example.com/confirmation",
				ValidationURL:   "invalid-url", // Invalid URL
			},
			wantErr: true,
		},

		// 9. Missing ConfirmationURL
		{
			name: "Missing ConfirmationURL",
			req: RegisterC2BURLRequest{
				ShortCode:     "12345",
				ResponseType:  types.ResponseType("SUCCESS"),
				CommandID:     types.CommandId("RegisterC2B"),
				ValidationURL: "https://example.com/validation",
			},
			wantErr: true,
		},

		// 10. Missing ValidationURL
		{
			name: "Missing ValidationURL",
			req: RegisterC2BURLRequest{
				ShortCode:       "12345",
				ResponseType:    types.ResponseType("SUCCESS"),
				CommandID:       types.CommandId("RegisterC2B"),
				ConfirmationURL: "https://example.com/confirmation",
			},
			wantErr: true,
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate the RegisterC2BURLRequest
			err := validate.Struct(tt.req)

			// Check if we expect an error
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err != nil)
			}
		})
	}
}

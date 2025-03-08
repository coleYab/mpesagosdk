package c2b

import (
	"testing"

	"github.com/coleYab/mpesagosdk/types"
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
				ShortCode:       "802000",
				ResponseType:    types.ResponseType("Completed"),
				CommandID:       types.CommandId("RegisterURL"),
				ConfirmationURL: "https://www.myservice:8080/confirmation",
				ValidationURL:   "https://www.myservice:8080/validation",
			},
			wantErr: false,
		},

		// 2. Missing ShortCode
		{
			name: "Missing ShortCode",
			req: RegisterC2BURLRequest{
				ResponseType:    types.ResponseType("Completed"),
				CommandID:       types.CommandId("RegisterURL"),
				ConfirmationURL: "https://www.myservice:8080/confirmation",
				ValidationURL:   "https://www.myservice:8080/validation",
			},
			wantErr: true,
		},

		// 3. ShortCode Too Short
		{
			name: "ShortCode Too Short",
			req: RegisterC2BURLRequest{
				ShortCode:       "802", // Too short, should be at least 5 characters
				ResponseType:    types.ResponseType("Completed"),
				CommandID:       types.CommandId("RegisterURL"),
				ConfirmationURL: "https://www.myservice:8080/confirmation",
				ValidationURL:   "https://www.myservice:8080/validation",
			},
			wantErr: true,
		},

		// 4. ShortCode Too Long
		{
			name: "ShortCode Too Long",
			req: RegisterC2BURLRequest{
				ShortCode:       "80200012345678901234567890", // Too long, should be at most 20 characters
				ResponseType:    types.ResponseType("Completed"),
				CommandID:       types.CommandId("RegisterURL"),
				ConfirmationURL: "https://www.myservice:8080/confirmation",
				ValidationURL:   "https://www.myservice:8080/validation",
			},
			wantErr: true,
		},

		// 5. Missing ResponseType
		{
			name: "Missing ResponseType",
			req: RegisterC2BURLRequest{
				ShortCode:       "802000",
				CommandID:       types.CommandId("RegisterURL"),
				ConfirmationURL: "https://www.myservice:8080/confirmation",
				ValidationURL:   "https://www.myservice:8080/validation",
			},
			wantErr: true,
		},

		// 6. Missing CommandID
		{
			name: "Missing CommandID",
			req: RegisterC2BURLRequest{
				ShortCode:       "802000",
				ResponseType:    types.ResponseType("Completed"),
				ConfirmationURL: "https://www.myservice:8080/confirmation",
				ValidationURL:   "https://www.myservice:8080/validation",
			},
			wantErr: true,
		},

		// 7. Invalid ConfirmationURL
		{
			name: "Invalid ConfirmationURL",
			req: RegisterC2BURLRequest{
				ShortCode:       "802000",
				ResponseType:    types.ResponseType("Completed"),
				CommandID:       types.CommandId("RegisterURL"),
				ConfirmationURL: "invalid-url", // Invalid URL
				ValidationURL:   "https://www.myservice:8080/validation",
			},
			wantErr: true,
		},

		// 8. Invalid ValidationURL
		{
			name: "Invalid ValidationURL",
			req: RegisterC2BURLRequest{
				ShortCode:       "802000",
				ResponseType:    types.ResponseType("Completed"),
				CommandID:       types.CommandId("RegisterURL"),
				ConfirmationURL: "https://www.myservice:8080/confirmation",
				ValidationURL:   "invalid-url", // Invalid URL
			},
			wantErr: true,
		},

		// 9. Missing ConfirmationURL
		{
			name: "Missing ConfirmationURL",
			req: RegisterC2BURLRequest{
				ShortCode:     "802000",
				ResponseType:  types.ResponseType("Completed"),
				CommandID:     types.CommandId("RegisterURL"),
				ValidationURL: "https://www.myservice:8080/validation",
			},
			wantErr: true,
		},

		// 10. Missing ValidationURL
		{
			name: "Missing ValidationURL",
			req: RegisterC2BURLRequest{
				ShortCode:       "802000",
				ResponseType:    types.ResponseType("Completed"),
				CommandID:       types.CommandId("RegisterURL"),
				ConfirmationURL: "https://www.myservice:8080/confirmation",
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

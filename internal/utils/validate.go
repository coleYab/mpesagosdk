package utils

import "github.com/go-playground/validator/v10"

// Validate performs struct-level validation using the provided validator instance.
//
// It takes a validator from the `go-playground/validator` package and a struct `data`
// to be validated. If the struct fails validation, it returns a `ValidationErrors` type
// containing all the validation errors. If the error is of another kind, it returns that error directly.
// If the struct passes validation, it returns nil.
//
// Parameters:
//	- v: an initialized *validator.Validate instance.
//	- data: the struct to be validated.
//
// Returns:
//	- error: nil if the struct is valid, otherwise a detailed validation error.
func Validate(v *validator.Validate, data any) error {
	if err := v.Struct(data); err != nil {
		errCasted, ok := err.(validator.ValidationErrors)
		if ok {
			return errCasted
		}
		return err
	}

	return nil
}

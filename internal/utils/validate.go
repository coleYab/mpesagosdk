package utils

import "github.com/go-playground/validator/v10"

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

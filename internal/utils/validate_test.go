package utils

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"required,gt=18"`
}

func TestValidate_Success(t *testing.T) {
	validate := validator.New()

	data := TestStruct{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   25,
	}

	err := Validate(validate, data)

	assert.NoError(t, err)
}

func TestValidate_InvalidName(t *testing.T) {
	validate := validator.New()

	data := TestStruct{
		Name:  "",
		Email: "john@example.com",
		Age:   25,
	}

	err := Validate(validate, data)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Name")
}

func TestValidate_InvalidEmail(t *testing.T) {
	validate := validator.New()

	data := TestStruct{
		Name:  "John Doe",
		Email: "notanemail",
		Age:   25,
	}

	err := Validate(validate, data)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Email")
}

func TestValidate_InvalidAge(t *testing.T) {
	validate := validator.New()

	data := TestStruct{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   16,
	}

	err := Validate(validate, data)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Age")
}

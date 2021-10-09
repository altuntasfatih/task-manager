package validator

import "github.com/go-playground/validator/v10"

func ValidateRequest(request interface{}) validator.ValidationErrors {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

package validator

import (
	"github.com/altuntasfatih/task-manager/pkg/models"
	"github.com/go-playground/validator/v10"
)

var customValidator = validator.New()

func createTaskRequestValidate(sl validator.StructLevel) {
	request := sl.Current().Interface().(models.CreateTaskRequest)
	if !request.StartTime.Before(request.EndTime) {
		sl.ReportError(sl.Current().Interface(), "CreateTaskRequest", "startTime", "", "startTime should be before endTime")
	}
}

func ValidateRequest(request interface{}) validator.ValidationErrors {

	customValidator.RegisterStructValidation(createTaskRequestValidate, models.CreateTaskRequest{})
	err := customValidator.Struct(request)
	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

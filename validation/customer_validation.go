package validation

import (
	"go-delivery-food/exception"
	model_request "go-delivery-food/model/request"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CreateCustomerValidate(request model_request.CreateCustomerRequest) {
	err := validation.ValidateStruct(&request, 
			validation.Field(&request.Id, validation.Required),
			validation.Field(&request.FirstName, validation.Required),
			validation.Field(&request.LastName, validation.Required),
			validation.Field(&request.Email, validation.Required),
			validation.Field(&request.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UpdateCustomerValidate(request model_request.UpdateCustomerRequest) {
	err := validation.ValidateStruct(&request, 
			validation.Field(&request.Id, validation.Required),
			validation.Field(&request.FirstName, validation.Required),
			validation.Field(&request.LastName, validation.Required),
			validation.Field(&request.Email, validation.Required),
			validation.Field(&request.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
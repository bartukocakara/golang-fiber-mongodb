package validation

import (
	"go-delivery-food/exception"
	model_request "go-delivery-food/model/request"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CreateMenuValidate(request model_request.CreateMenuRequest) {
	err := validation.ValidateStruct(&request, 
			validation.Field(&request.Id, validation.Required),
			validation.Field(&request.Name, validation.Required),
			validation.Field(&request.Price, validation.Required, validation.Min(1)),
			validation.Field(&request.Quantity, validation.Required, validation.Min(0)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UpdateMenuValidate(request model_request.UpdateMenuRequest) {
	err := validation.ValidateStruct(&request, 
			validation.Field(&request.Id, validation.Required),
			validation.Field(&request.Name, validation.Required),
			validation.Field(&request.Price, validation.Required, validation.Min(1)),
			validation.Field(&request.Quantity, validation.Required, validation.Min(0)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
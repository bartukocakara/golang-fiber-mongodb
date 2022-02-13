package customer_service

import (
	model_request "go-delivery-food/model/request"
	model_response "go-delivery-food/model/response"
)

type CustomerService interface {
	Create(request model_request.CreateCustomerRequest) (response model_response.CreateCustomerResponse)
	List(response model_response.ListCustomerResponse)
	Update(response model_response.UpdateCustomerResponse)
}
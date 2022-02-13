package controller

import "go-delivery-food/service/customer_service"

type CustomerController struct {
	CustomerService customer_service.CustomerService
}
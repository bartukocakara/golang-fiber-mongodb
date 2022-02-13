package customer_repository

import "go-delivery-food/entity"

type CustomerRepository interface {
	Insert(customer entity.Customer)

	List() (customers []entity.Customer)

	Delete()

	DeleteAll()
}
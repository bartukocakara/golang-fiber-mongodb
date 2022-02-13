package customer_repository

import (
	"go-delivery-food/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type customerRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewCustomerRepository(database *mongo.Database) CustomerRepository {
	return &customerRepositoryImpl{
		Collection : database.Collection("customers"),
	}
}


func (repository *customerRepositoryImpl) List() (customers []entity.Customer) {
	
}


func (repository *customerRepositoryImpl) Insert(customer entity.Customer) {

}

func (repository *customerRepositoryImpl) Delete() () {
	
}

func (repository *customerRepositoryImpl) DeleteAll() {
	
}

package menu_repository

import (
	"go-delivery-food/config"
	"go-delivery-food/entity"
	"go-delivery-food/exception"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type menuRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewMenuRepository(database *mongo.Database) MenuRepository {
	return &menuRepositoryImpl{
		Collection: database.Collection("menus"),
	}
}

func (repository *menuRepositoryImpl) List()(menus []entity.Menu){
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})

	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)

	exception.PanicIfNeeded(err)

	for _, document := range documents {
		menus = append(menus,	entity.Menu{
			Id : document["_id"].(string),
			Name : document["name"].(string),
			Price : document["price"].(int64),
			Quantity : document["quantity"].(int32),
		})
	}

	return menus
}


func (repository *menuRepositoryImpl) Insert(menu entity.Menu) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M {
		"_id" : menu.Id,
		"name": menu.Name,
		"price": menu.Price,
		"quantity" : menu.Quantity,
	})
	exception.PanicIfNeeded(err)
}

func (repository *menuRepositoryImpl) Delete(){

}

func (repository *menuRepositoryImpl) DeleteAll(){
	ctx, cancel := config.NewMongoContext()

	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}


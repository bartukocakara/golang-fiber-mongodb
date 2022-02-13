package menu_item_repository

import (
	"go-delivery-food/config"
	"go-delivery-food/entity"
	"go-delivery-food/exception"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMenuItemRepository(database *mongo.Database) MenuItemRepository {
	return &menuItemRepositoryImpl{
		Collection: database.Collection("menu_items"),
	}
}

type menuItemRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *menuItemRepositoryImpl) List()(menuItems []entity.MenuItem){
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})

	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)

	exception.PanicIfNeeded(err)

	for _, document := range documents {
		menuItems = append(menuItems,	entity.MenuItem{
			Id : document["_id"].(string),
			Name : document["name"].(string),
			Price : document["price"].(int64),
			Quantity : document["quantity"].(int32),
		})
	}

	return menuItems
}


func (repository *menuItemRepositoryImpl) Insert(menuItem entity.MenuItem) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M {
		"_id" : menuItem.Id,
		"name": menuItem.Name,
		"price": menuItem.Price,
		"quantity" : menuItem.Quantity,
	})
	exception.PanicIfNeeded(err)
}

func (repository *menuItemRepositoryImpl) Delete(){

}

func (repository *menuItemRepositoryImpl) DeleteAll(){
	ctx, cancel := config.NewMongoContext()

	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}


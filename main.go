package main

import (
	"go-delivery-food/config"
	"go-delivery-food/controller"
	"go-delivery-food/exception"
	"go-delivery-food/repository"
	"go-delivery-food/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	menuRepository := repository.NewMenuRepository(database)

	menuService := service.NewMenuService(&menuRepository)

	menuController := controller.NewMenuController(&menuService)


	app:= fiber.New(config.NewFiberConfig())

	app.Use(recover.New())

	menuController.Route(app)

	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)


}
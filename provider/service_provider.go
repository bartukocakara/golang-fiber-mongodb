package provider

import (
	"go-delivery-food/config"
	"go-delivery-food/controller"
	"go-delivery-food/repository/menu_item_repository"
	"go-delivery-food/repository/menu_repository"
	"go-delivery-food/service/menu_item_service"
	"go-delivery-food/service/menu_service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func ProvideService()(app *fiber.App) {
	
	app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	configuration := config.New()
	database := config.NewMongoDatabase(configuration)
	// urlPrefix := config.Version()
	menuRepository := menu_repository.NewMenuRepository(database)
	menuService := menu_service.NewMenuService(&menuRepository)
	menuController := controller.NewMenuController(&menuService)

	menuItemRepository := menu_item_repository.NewMenuItemRepository(database)
	menuItemService := menu_item_service.NewMenuItemService(&menuItemRepository)
	menuItemController := controller.NewMenuItemController(&menuItemService)

	menuController.MenuRoute(app)
	menuItemController.MenuItemRoute(app)
	return app
}
package controller

import (
	"go-delivery-food/exception"
	"go-delivery-food/model"
	model_request "go-delivery-food/model/request"
	"go-delivery-food/service/menu_item_service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
type MenuItemController struct {
	MenuItemService menu_item_service.MenuItemService
}

func NewMenuItemController(menuItemService *menu_item_service.MenuItemService) MenuItemController {
	return MenuItemController{MenuItemService: *menuItemService}
}

func (controller *MenuItemController) MenuItemRoute(app *fiber.App){
	app.Post("api/menu-items", controller.CreateMenuItem)
	app.Get("api/menu-items", controller.ListMenuItem)
}

func (controller *MenuItemController) CreateMenuItem(c *fiber.Ctx) error {
	var request model_request.CreateMenuItemRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()
	exception.PanicIfNeeded(err)

	response := controller.MenuItemService.Create(request)
	return c.JSON(model.ApiResponse{
		Code : 200,
		Status : "OK",
		Data : response,
	})

}

func (controller *MenuItemController) ListMenuItem(c *fiber.Ctx) error {
	responses := controller.MenuItemService.List()
	return c.JSON(model.ApiResponse{
		Code : 200,
		Status : "OK",
		Data: responses,
	})
}
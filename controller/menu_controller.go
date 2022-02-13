package controller

import (
	"go-delivery-food/exception"
	model "go-delivery-food/model"
	model_request "go-delivery-food/model/request"
	"go-delivery-food/service/menu_service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuController struct {
	MenuService menu_service.MenuService
}

func NewMenuController(menuService *menu_service.MenuService) MenuController {
	return MenuController{MenuService: *menuService}
}

func (controller *MenuController) MenuRoute(app *fiber.App){
	app.Post("/api/menus", controller.CreateMenu)
	app.Get("api/menus", controller.ListMenu)
}

func (controller *MenuController) CreateMenu(c *fiber.Ctx) error {
	var request model_request.CreateMenuRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()
	exception.PanicIfNeeded(err)

	response := controller.MenuService.Create(request)
	return c.JSON(model.ApiResponse{
		Code : 200,
		Status: "OK",
		Data: response,
	})
}

func (controller *MenuController) ListMenu(c *fiber.Ctx) error {
	responses := controller.MenuService.List()
	return c.JSON(model.ApiResponse{
		Code: 200,
		Status: "OK",
		Data: responses,
	})
}
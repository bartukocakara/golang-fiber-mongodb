package controller

import (
	"go-delivery-food/exception"
	"go-delivery-food/model"
	"go-delivery-food/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MenuController struct {
	MenuService service.MenuService
}

func NewMenuController(menuService *service.MenuService) MenuController {
	return MenuController{MenuService: *menuService}
}

func (controller *MenuController) Route(app *fiber.App){
	app.Post("/api/menus", controller.Create)
	app.Get("api/menus", controller.List)
}

func (controller *MenuController) Create(c *fiber.Ctx) error {
	var request model.CreateMenuRequest
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

func (controller *MenuController) List(c *fiber.Ctx) error {
	responses := controller.MenuService.List()
	return c.JSON(model.ApiResponse{
		Code: 200,
		Status: "OK",
		Data: responses,
	})
}
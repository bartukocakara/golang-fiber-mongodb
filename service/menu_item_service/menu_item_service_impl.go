package menu_item_service

import (
	"go-delivery-food/entity"
	model_request "go-delivery-food/model/request"
	model_response "go-delivery-food/model/response"
	"go-delivery-food/repository/menu_item_repository"
	"go-delivery-food/validation"
)

func NewMenuItemService(menuItemRepository *menu_item_repository.MenuItemRepository) MenuItemService {
	return &menuItemServiceImpl{
		MenuItemRepository : *menuItemRepository,
	}
}

type menuItemServiceImpl struct {
	MenuItemRepository menu_item_repository.MenuItemRepository
}

func (service *menuItemServiceImpl) Create(request model_request.CreateMenuItemRequest) (response model_response.CreateMenuItemResponse) {
	validation.CreateMenuItemValidate(request)

	menuItem := entity.MenuItem {
		Id : request.Id,
		Name : request.Name,
		Price: request.Price,
		Quantity : request.Quantity,
	}
	service.MenuItemRepository.Insert(menuItem)

	response = model_response.CreateMenuItemResponse{
		Id : menuItem.Id,
		Name: menuItem.Name,
		Price: menuItem.Price,
		Quantity : menuItem.Quantity,
	}
	return response
}

func(service *menuItemServiceImpl) List() (responses []model_response.ListMenuItemsResponse){
	menuItems := service.MenuItemRepository.List()
	for _, menuItem := range menuItems {
		responses = append(responses, model_response.ListMenuItemsResponse{
			Id : menuItem.Id,
			Name : menuItem.Name,
			Price: menuItem.Price,
			Quantity : menuItem.Quantity,
		})
	}

	return responses
}

func(service *menuItemServiceImpl) Delete(){

}

func(service *menuItemServiceImpl) DeleteAll(){

}
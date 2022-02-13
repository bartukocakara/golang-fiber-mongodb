package menu_service

import (
	"go-delivery-food/entity"
	model_request "go-delivery-food/model/request"
	model_response "go-delivery-food/model/response"
	"go-delivery-food/repository/menu_repository"
	"go-delivery-food/validation"
)

func NewMenuService(menuRepository *menu_repository.MenuRepository) MenuService {
	return &menuServiceImpl{
		MenuRepository: *menuRepository,
	}
}

type menuServiceImpl struct {
	MenuRepository menu_repository.MenuRepository
}

func (service *menuServiceImpl) Create(request model_request.CreateMenuRequest) (response model_response.CreateMenuResponse) {
	validation.CreateMenuValidate(request)

	menu := entity.Menu {
		Id : request.Id,
		Name : request.Name,
		Price: request.Price,
		Quantity : request.Quantity,
	}
	service.MenuRepository.Insert(menu)

	response = model_response.CreateMenuResponse{
		Id : menu.Id,
		Name: menu.Name,
		Price: menu.Price,
		Quantity : menu.Quantity,
	}
	return response
}

func(service *menuServiceImpl) List() (responses []model_response.ListMenusResponse){
	menus := service.MenuRepository.List()
	for _, menu := range menus {
		responses = append(responses, model_response.ListMenusResponse{
			Id : menu.Id,
			Name : menu.Name,
			Price: menu.Price,
			Quantity : menu.Quantity,
		})
	}

	return responses
}

func(service *menuServiceImpl) Delete(){

}

func(service *menuServiceImpl) DeleteAll(){

}
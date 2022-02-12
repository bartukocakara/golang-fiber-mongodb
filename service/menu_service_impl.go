package service

import (
	"go-delivery-food/entity"
	"go-delivery-food/model"
	"go-delivery-food/repository"
	"go-delivery-food/validation"
)

func NewMenuService(menuRepository *repository.MenuRepository) MenuService {
	return &menuServiceImpl{
		MenuRepository: *menuRepository,
	}
}

type menuServiceImpl struct {
	MenuRepository repository.MenuRepository
}

func (service *menuServiceImpl) Create(request model.CreateMenuRequest) (response model.CreateMenuResponse) {
	validation.Validate(request)

	menu := entity.Menu {
		Id : request.Id,
		Name : request.Name,
		Price: request.Price,
		Quantity : request.Quantity,
	}
	service.MenuRepository.Insert(menu)

	response = model.CreateMenuResponse{
		Id : menu.Id,
		Name: menu.Name,
		Price: menu.Price,
		Quantity : menu.Quantity,
	}
	return response
}

func(service *menuServiceImpl) List() (responses []model.ListMenusResponse){
	menus := service.MenuRepository.List()
	for _, menu := range menus {
		responses = append(responses, model.ListMenusResponse{
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
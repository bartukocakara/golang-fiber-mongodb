package service

import "go-delivery-food/model"

type MenuService interface {
	Create(request model.CreateMenuRequest) (response model.CreateMenuResponse)
	List() (responses []model.ListMenusResponse)
	Delete()
	DeleteAll()
}
package menu_service

import (
	model_request "go-delivery-food/model/request"
	model_response "go-delivery-food/model/response"
)

type MenuService interface {
	Create(request model_request.CreateMenuRequest) (response model_response.CreateMenuResponse)
	List() (responses []model_response.ListMenusResponse)
	Delete()
	DeleteAll()
}
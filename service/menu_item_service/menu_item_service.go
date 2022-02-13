package menu_item_service

import (
	model_request "go-delivery-food/model/request"
	model_response "go-delivery-food/model/response"
)

type MenuItemService interface {
	Create(request model_request.CreateMenuItemRequest) (response model_response.CreateMenuItemResponse)
	List() (responses []model_response.ListMenuItemsResponse)
	Delete()
	DeleteAll()
}
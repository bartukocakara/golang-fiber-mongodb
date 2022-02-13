package menu_repository

import "go-delivery-food/entity"

type MenuRepository interface {
	Insert(menu entity.Menu)

	List() (menus []entity.Menu)

	Delete()

	DeleteAll()
}
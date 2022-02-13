package menu_item_repository

import "go-delivery-food/entity"

type MenuItemRepository interface {
	Insert(menuItem entity.MenuItem)

	List() (menuItems []entity.MenuItem)

	Delete()

	DeleteAll()
}
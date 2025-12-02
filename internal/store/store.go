package store

import "Order5003/internal/models"

type Store interface {
	GetAllMenuItems() []models.Menu
	GetMenuItemByID(id int) (models.Menu, error)
	CreateMenuItem(item models.Menu) models.Menu
	UpdateMenuItem(id int, item models.Menu) (models.Menu, error)
	DeleteMenuItem(id int) error
	CreateOrder(order models.Order) models.Order
	GetOrderByID(id int) (models.Order, error)
	GetAllOrders() []models.Order
	UpdateOrderStatus(id int, status models.OrderStatus) (models.Order, error)
	GetUserByUsername(username string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	GetShopByName(name string) (models.Shop, error)
	GetDelivererByName(name string) (models.Deliverers, error)
	GetRandomTableNumber() string
}

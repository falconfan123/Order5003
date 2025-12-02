package store

import "Order5003/internal/models"

type Store interface {
	GetAllMenuItems() []models.MenuItem
	GetMenuItemByID(id int) (models.MenuItem, error)
	CreateMenuItem(item models.MenuItem) models.MenuItem
	UpdateMenuItem(id int, item models.MenuItem) (models.MenuItem, error)
	DeleteMenuItem(id int) error
	CreateOrder(order models.Order) models.Order
	GetOrderByID(id int) (models.Order, error)
	GetAllOrders() []models.Order
	UpdateOrderStatus(id int, status models.OrderStatus) (models.Order, error)
	GetUserByUsername(username string) (models.User, error)
	GetShopByName(name string) (models.Shop, error)
	GetRandomTableNumber() string
}

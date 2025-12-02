package service

import "Order5003/internal/bizmodel"

type MenuService interface {
	GetAllMenuItems() []bizmodel.Menu
	GetMenuItemByID(id int) (bizmodel.Menu, error)
	CreateMenuItem(item bizmodel.Menu) bizmodel.Menu
	UpdateMenuItem(id int, item bizmodel.Menu) (bizmodel.Menu, error)
	DeleteMenuItem(id int) error
}

type OrderService interface {
	CreateOrder(order bizmodel.Order) bizmodel.Order
	GetOrderByID(id int) (bizmodel.Order, error)
	GetAllOrders() []bizmodel.Order
	UpdateOrderStatus(id int, status bizmodel.OrderStatus) (bizmodel.Order, error)
}

type ShopService interface {
	GetShopByName(name string) (bizmodel.Shop, error)
	GetAllShops() ([]bizmodel.Shop, error)
	GetAllMenuDishesByShopID(shopID int) ([]bizmodel.Dishes, error)
}

type UserService interface {
	GetUserByUsername(username string) (bizmodel.User, error)
	GetUserByID(id int) (bizmodel.User, error)
}

type DelivererService interface {
	GetDelivererByName(name string) (bizmodel.Deliverers, error)
}

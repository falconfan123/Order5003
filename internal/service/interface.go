package service

import (
	"Order5003/internal/bizmodel"
	"context"

	"gorm.io/gorm"
)

type MenuService interface {
	GetAllMenuItems() []bizmodel.Menu
	GetMenuItemByID(id int) (bizmodel.Menu, error)
	UpdateMenuItem(id int, item bizmodel.Menu) (bizmodel.Menu, error)
	DeleteMenuItem(id int) error
}

type OrderService interface {
	CreateOrder(order bizmodel.Order) bizmodel.Order
	GetOrderByID(id int) (bizmodel.Order, error)
	GetAllOrders() []bizmodel.Order
	UpdateOrderStatus(id int, status bizmodel.OrderStatus) (bizmodel.Order, error)
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	GetDishByID(ctx context.Context, tx *gorm.DB, dishID int) (*bizmodel.Dishes, error)
	CreateOrderMaster(ctx context.Context, tx *gorm.DB, orderMaster *bizmodel.Order) (int, error)
	CreateOrderDish(ctx context.Context, tx *gorm.DB, orderDish *bizmodel.OrderDishDetail) error
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

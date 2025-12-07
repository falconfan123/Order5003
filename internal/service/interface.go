package service

import (
	"Order5003/internal/bizmodel"
	"context"

	"gorm.io/gorm"
)

type MenuService interface {
	GetMenuByShopID(ctx context.Context, shopID int) ([]bizmodel.Menu, error)
	GetDishesByMenuID(ctx context.Context, menuID int) ([]bizmodel.Dishes, error)
}

type OrderService interface {
	CreateOrder(order bizmodel.Order) bizmodel.Order
	GetOrderByUserID(userID int) ([]bizmodel.Order, error)
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	GetDishByID(ctx context.Context, tx *gorm.DB, dishID int) (*bizmodel.Dishes, error)
	CreateOrderMaster(ctx context.Context, tx *gorm.DB, orderMaster *bizmodel.Order) (int, error)
	CreateOrderDish(ctx context.Context, tx *gorm.DB, orderDish *bizmodel.OrderDishDetail) error
	GetOrderDishesByOrderID(ctx context.Context, orderID int) ([]bizmodel.OrderDishDetail, error)
}

type ShopService interface {
	GetOrdersByShopID(shopID int) ([]bizmodel.Order, error)
	GetShopByName(name string) (bizmodel.Shop, error)
	GetAllShops() ([]bizmodel.Shop, error)
	GetMenuDishesByShopID(shopID int) ([]bizmodel.Dishes, error)
	GetShopNameByShopID(shopID int) (string, error)
	GetDeliveryFeeByShopID(shopID int) (float64, error)
	GetBusinessHoursByShopID(shopID int) (string, error)
	GetShopTypeByShopID(shopID int) (int, error)
	//查看店铺状态
	GetShopStatusByShopID(shopID int) (int, error)
	UpdateShopStatus(shopID int, status int) (int, error)
	GetDishesByOrderID(orderID int) ([]bizmodel.Dishes, error)
	AcceptOrder(orderID int) error
	WaitingForDeliveryOrder(orderID int) error
}

type UserService interface {
	GetUserByUsername(username string) (bizmodel.User, error)
	GetUserByID(id int) (bizmodel.User, error)
}

type DelivererService interface {
	GetDelivererByName(name string) (bizmodel.Deliverers, error)
}

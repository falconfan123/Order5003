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
	//停售菜品
	StopDish(dishID int) error
	//上架商品
	StartDish(dishID int) error
	//今日订单数
	GetTodayOrderCountByShopID(shopID int) (int, error)
	//写入今日订单数
	WriteTodayOrderCountByShopID(shopID int, count int) error
	//获取今日营业额
	GetTodayRevenueByShopID(shopID int) (float64, error)
	//写入今日营业额
	WriteTodayRevenueByShopID(shopID int, revenue float64) error
	//获取历史所有订单数
	GetAllOrderCountByShopID(shopID int) ([]bizmodel.OrderCountAll, error)
	//获取历史所有营业额
	GetAllRevenueByShopID(shopID int) ([]bizmodel.RevenueCountAll, error)
}

type UserService interface {
	GetUserByUsername(username string) (bizmodel.User, error)
	GetUserByID(id int) (bizmodel.User, error)
	GetUsernameByUserID(userID int) (string, error)
	GetUserAddressByUserID(userID int) (string, error)
}

type DelivererService interface {
	GetDelivererByName(name string) (bizmodel.Deliverers, error)
}

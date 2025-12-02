package store

import "Order5003/internal/bizmodel"

type Store interface {
    GetAllMenuItems() []bizmodel.Menu
    GetMenuItemByID(id int) (bizmodel.Menu, error)
    CreateMenuItem(item bizmodel.Menu) bizmodel.Menu
    UpdateMenuItem(id int, item bizmodel.Menu) (bizmodel.Menu, error)
    DeleteMenuItem(id int) error
    CreateOrder(order bizmodel.Order) bizmodel.Order
    GetOrderByID(id int) (bizmodel.Order, error)
    GetAllOrders() []bizmodel.Order
    UpdateOrderStatus(id int, status bizmodel.OrderStatus) (bizmodel.Order, error)
    GetUserByUsername(username string) (bizmodel.User, error)
    GetUserByID(id int) (bizmodel.User, error)
    GetShopByName(name string) (bizmodel.Shop, error)
    GetDelivererByName(name string) (bizmodel.Deliverers, error)
    GetAllShops() ([]bizmodel.Shop, error)
}

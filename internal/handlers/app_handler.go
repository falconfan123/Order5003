package handlers

import "Order5003/internal/service"

type AppHandlers struct {
    Menu     *MenuHandler
    Order    *OrderHandler
    User     *UserHandler
    Shop     *ShopHandler
    Deliver  *DeliverHandler
    MenuDish *MenuDishHandler
}

func NewAppHandlers(menu service.MenuService, order service.OrderService, user service.UserService, shop service.ShopService, deliver service.DelivererService) *AppHandlers {
    return &AppHandlers{
        Menu:     NewMenuHandler(menu),
        Order:    NewOrderHandler(order),
        User:     NewUserHandler(user),
        Shop:     NewShopHandler(shop),
        Deliver:  NewDeliverHandler(deliver),
        MenuDish: NewMenuDishHandler(shop),
    }
}

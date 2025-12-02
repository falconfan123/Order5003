package handlers

import "Order5003/internal/store"

type AppHandlers struct {
    Menu    *MenuHandler
    Order   *OrderHandler
    User    *UserHandler
    Shop    *ShopHandler
    Deliver *DeliverHandler
}

func NewAppHandlers(s store.Store) *AppHandlers {
    return &AppHandlers{
        Menu:    NewMenuHandler(s),
        Order:   NewOrderHandler(s),
        User:    NewUserHandler(s),
        Shop:    NewShopHandler(s),
        Deliver: NewDeliverHandler(s),
    }
}


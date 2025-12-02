package impl

import (
    "Order5003/internal/bizmodel"
    "Order5003/internal/dao"
    "Order5003/internal/model"
    "errors"
)

func (s *GormStore) CreateOrder(order bizmodel.Order) bizmodel.Order {
    e := &model.OrderEntity{}
    if _, err := dao.CreateOrder(s.db, e); err != nil {
        return order
    }
    order.ID = e.OrderID
    return order
}

func (s *GormStore) GetOrderByID(id int) (bizmodel.Order, error) {
    e, err := dao.GetOrderByID(s.db, id)
    if err != nil {
        return bizmodel.Order{}, errors.New("order not found")
    }
    var items []bizmodel.OrderItem
    return bizmodel.Order{
        ID:        e.OrderID,
        Items:     items,
        Status:    bizmodel.OrderStatus(e.Status),
        CreatedAt: e.CreatedAt,
    }, nil
}

func (s *GormStore) GetAllOrders() []bizmodel.Order {
    list, err := dao.ListOrders(s.db)
    if err != nil {
        return []bizmodel.Order{}
    }
    out := make([]bizmodel.Order, 0, len(list))
    for _, e := range list {
        var items []bizmodel.OrderItem
        out = append(out, bizmodel.Order{
            ID:        e.OrderID,
            Items:     items,
            Status:    bizmodel.OrderStatus(e.Status),
            CreatedAt: e.CreatedAt,
        })
    }
    return out
}

func (s *GormStore) UpdateOrderStatus(id int, status bizmodel.OrderStatus) (bizmodel.Order, error) {
    if err := dao.UpdateOrderStatus(s.db, id, string(status)); err != nil {
        return bizmodel.Order{}, err
    }
    return s.GetOrderByID(id)
}

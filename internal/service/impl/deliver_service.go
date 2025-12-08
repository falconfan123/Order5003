package impl

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"Order5003/internal/model"
	"errors"
)

func (s *GormStore) GetDelivererByName(name string) (bizmodel.Deliverers, error) {
	e, err := dao.GetDelivererByName(s.db, name)
	if err != nil {
		return bizmodel.Deliverers{}, errors.New("deliverer not found")
	}
	return bizmodel.Deliverers{
		DelivererID:     e.DelivererID,
		Name:            e.Name,
		Phone:           e.Phone,
		ResponsibleArea: e.ResponsibleArea,
		Password:        e.Password,
	}, nil
}

func (s *GormStore) GetOrderWaitingForDeliver() ([]bizmodel.Order, error) {
	orders, err := dao.GetOrderWaitingForDeliver(s.db)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	out := make([]bizmodel.Order, 0, len(orders))
	for _, order := range orders {
		out = append(out, bizmodel.Order{
			OrderID:     order.OrderID,
			UserID:      order.UserID,
			ShopID:      order.ShopID,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
			CreatedAt:   order.CreatedAt,
		})
	}
	return out, nil
}

func (s *GormStore) AcceptOrderDeliver(deliverID int, orderID int) error {
	err := dao.AcceptOrder(s.db, deliverID, orderID)
	//改变订单状态为配送中
	err = dao.UpdateOrderStatus(s.db, orderID, int(model.OrderStatusDelivering))
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}

func (s *GormStore) GetMyOrder(deliverID int) ([]bizmodel.Order, error) {
	delivery, err := dao.GetMyOrder(s.db, deliverID)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	out := make([]bizmodel.Order, 0, len(delivery))
	for _, d := range delivery {
		order, err := dao.GetOrderByOrderID(s.db, d.OrderID)
		if err != nil {
			return nil, errors.New("internal server error")
		}
		out = append(out, bizmodel.Order{
			OrderID:     order.OrderID,
			UserID:      order.UserID,
			ShopID:      order.ShopID,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
			CreatedAt:   order.CreatedAt,
		})
	}
	return out, nil
}

func (s *GormStore) ConfirmDeliver(deliverID int, orderID int) error {
	//更改delivery表下的信息
	err := dao.ConfirmDeliver(s.db, deliverID, orderID)
	//更改order表下的信息
	err = dao.UpdateOrderStatus(s.db, orderID, int(model.OrderStatusCompleted))
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}

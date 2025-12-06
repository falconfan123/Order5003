package impl

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"Order5003/internal/logger"
	"errors"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func (s *GormStore) GetShopByName(name string) (bizmodel.Shop, error) {
	e, err := dao.GetShopByName(s.db, name)
	if err != nil {
		return bizmodel.Shop{}, errors.New("shop not found")
	}
	return bizmodel.Shop{
		ShopID:        e.ShopID,
		ShopName:      e.ShopName,
		DeliveryRange: e.DeliveryRange,
		DeliveryFee:   e.DeliveryFee,
		BusinessHours: e.BusinessHours,
		CreatedAt:     &e.CreatedAt,
		Password:      e.Password,
	}, nil
}

func (s *GormStore) GetAllShops() ([]bizmodel.Shop, error) {
	list, err := dao.ListShops(s.db)
	logger.Info("list shops", zap.Any("list", list))
	if err != nil {
		return []bizmodel.Shop{}, errors.New("shops not found")
	}
	out := make([]bizmodel.Shop, 0, len(list))
	for _, e := range list {
		out = append(out, bizmodel.Shop{
			ShopID:        e.ShopID,
			ShopName:      e.ShopName,
			DeliveryRange: e.DeliveryRange,
			DeliveryFee:   e.DeliveryFee,
			BusinessHours: e.BusinessHours,
			CreatedAt:     &e.CreatedAt,
			Password:      e.Password,
			Status:        int(e.Status),
			Type:          int(e.Type),
		})
	}
	return out, nil
}

func (s *GormStore) GetShopNameByShopID(shopID int) (string, error) {
	e, err := dao.GetShopByID(s.db, shopID)
	if err != nil {
		return "", errors.New("shop not found")
	}
	return e.ShopName, nil
}

// GetOrdersByShopID 获取指定店铺的所有订单
func (s *GormStore) GetOrdersByShopID(shopID int) ([]bizmodel.Order, error) {
	list, err := dao.ListOrdersByShopID(s.db, shopID)
	logger.Info("list orders by shop id", zap.Any("list", list))
	if err != nil {
		return []bizmodel.Order{}, errors.New("orders not found")
	}
	out := make([]bizmodel.Order, 0, len(list))
	for _, e := range list { //list: []model.OrderEntity
		out = append(out, bizmodel.Order{
			OrderID:     e.OrderID,
			ShopID:      e.ShopID,
			UserID:      e.UserID,
			TotalAmount: e.TotalAmount,
			Status:      int(e.Status),
			CreatedAt:   e.CreatedAt,
		})
	}
	return out, nil
}

// GetDeliveryFeeByShopID 获取指定店铺的配送费
func (s *GormStore) GetDeliveryFeeByShopID(shopID int) (decimal.Decimal, error) {
	e, err := dao.GetShopByID(s.db, shopID)
	if err != nil {
		return decimal.Decimal{}, errors.New("shop not found")
	}
	return e.DeliveryFee, nil
}

// GetBusinessHoursByShopID 获取指定店铺的营业时间
func (s *GormStore) GetBusinessHoursByShopID(shopID int) (string, error) {
	e, err := dao.GetShopByID(s.db, shopID)
	if err != nil {
		return "", errors.New("shop not found")
	}
	return e.BusinessHours, nil
}

// GetShopTypeByShopID 获取指定店铺的类型
func (s *GormStore) GetShopTypeByShopID(shopID int) (int, error) {
	e, err := dao.GetShopByID(s.db, shopID)
	if err != nil {
		return 0, errors.New("shop not found")
	}
	return int(e.Status), nil
}

// UpdateShopStatus 更新指定店铺的状态
func (s *GormStore) UpdateShopStatus(shopID int, status int) (int, error) {
	e, err := dao.GetShopByID(s.db, shopID)
	if err != nil {
		return 0, errors.New("shop not found")
	}
	if err := dao.UpdateShopStatus(s.db, e, status); err != nil {
		return 0, errors.New("update shop failed")
	}
	return int(e.Status), nil
}

func (s *GormStore) GetDishesByOrderID(orderID int) ([]bizmodel.Dishes, error) {
	OrderDish, err := dao.GetOrderDishesByOrderID(s.db, orderID)
	logger.Info("list order dishes by order id", zap.Any("list", OrderDish), zap.Int("orderID", orderID))
	if err != nil {
		return []bizmodel.Dishes{}, errors.New("dishes not found")
	}
	out := make([]bizmodel.Dishes, 0, len(OrderDish))
	for _, e := range OrderDish {
		dish, err := dao.GetDishByDishID(s.db, e.DishID)
		if err != nil {
			return []bizmodel.Dishes{}, errors.New("dishes not found")
		}
		out = append(out, bizmodel.Dishes{
			DishID:   dish.DishID,
			ShopID:   dish.ShopID,
			DishName: dish.DishName,
			Price:    dish.Price,
			Status:   int(dish.Status),
		})
	}
	return out, nil
}

// AcceptOrder 接受指定订单
func (s *GormStore) AcceptOrder(orderID int) error {
	if err := dao.UpdateOrderStatus(s.db, orderID, int(bizmodel.OrderStatusPreparing)); err != nil {
		return errors.New("update order status failed")
	}
	return nil
}

// WaitingForDeliveryOrder 订单配送中
func (s *GormStore) WaitingForDeliveryOrder(orderID int) error {
	if err := dao.UpdateOrderStatus(s.db, orderID, int(bizmodel.OrderStatusWaitingForDelivery)); err != nil {
		return errors.New("update order status failed")
	}
	return nil
}

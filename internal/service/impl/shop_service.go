package impl

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"Order5003/internal/logger"
	"Order5003/internal/model"
	"errors"
	"time"

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
		DeliveryRange: e.DeliveryRange.InexactFloat64(),
		DeliveryFee:   e.DeliveryFee.InexactFloat64(),
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
			DeliveryRange: e.DeliveryRange.InexactFloat64(),
			DeliveryFee:   e.DeliveryFee.InexactFloat64(),
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
func (s *GormStore) GetDeliveryFeeByShopID(shopID int) (float64, error) {
	e, err := dao.GetShopByID(s.db, shopID)
	if err != nil {
		return 0, errors.New("shop not found")
	}
	return e.DeliveryFee.InexactFloat64(), nil
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

// GetShopStatusByShopID 获取指定店铺的状态
func (s *GormStore) GetShopStatusByShopID(shopID int) (int, error) {
	e, err := dao.GetShopByID(s.db, shopID)
	if err != nil {
		return 0, errors.New("shop not found")
	}
	return int(e.Status), nil
}

// StopDish 停售指定菜品
func (s *GormStore) StopDish(dishID int) error {
	if err := dao.UpdateDishStatus(s.db, dishID, int(bizmodel.DishStatusStopped)); err != nil {
		return errors.New("update dish status failed")
	}
	return nil
}

// StartDish 上架指定菜品
func (s *GormStore) StartDish(dishID int) error {
	if err := dao.UpdateDishStatus(s.db, dishID, int(bizmodel.DishStatusAvailable)); err != nil {
		return errors.New("update dish status failed")
	}
	return nil
}

// GetTodayOrderCountByShopID 获取指定店铺的今日订单数
func (s *GormStore) GetTodayOrderCountByShopID(shopID int) (int, error) {
	count, err := dao.GetTodayOrderCountByShopID(s.db, shopID)
	if err != nil {
		return 0, errors.New("get today order count failed")
	}
	return count, nil
}

// WriteTodayOrderCountByShopID 写入指定店铺的今日订单数
func (s *GormStore) WriteTodayOrderCountByShopID(shopID int, count int) error {
	//先看是否已经存在 如果已经存在 则只变更count字段
	existing, err := dao.GetShopDailyRevenueByShopIDAndDate(s.db, shopID, time.Now().Truncate(24*time.Hour))
	if err != nil {
		return errors.New("get shop daily revenue failed")
	}
	logger.Info("existing", zap.Any("existing", existing.ID))
	if existing.ID > 0 {
		logger.Info("count", zap.Int("count", count))
		if err := dao.UpdateShopDailyRevenueCount(s.db, existing, count); err != nil {
			return errors.New("update shop daily revenue count failed")
		}
		return nil
	}
	//如果不存在 则插入新记录
	if err := dao.CreateShopDailyRevenue(s.db, model.ShopDailyRevenue{
		ShopID:     int64(shopID),
		Date:       time.Now().Truncate(24 * time.Hour),
		OrderCount: count,
	}); err != nil {
		return errors.New("create shop daily revenue failed")
	}
	return nil
}

// GetTodayRevenueByShopID 获取指定店铺的今日营业额
func (s *GormStore) GetTodayRevenueByShopID(shopID int) (float64, error) {
	//统计今日营业额 也就是status=8 且data是今日的orders总和
	TodayFinishOrder, err := dao.GetTodayFinishOrderByShopID(s.db, shopID)
	if err != nil {
		return 0, errors.New("get today finish order failed")
	}
	revenue := 0.0
	for _, order := range TodayFinishOrder {
		revenue += order.TotalAmount.InexactFloat64()
	}
	//将更新的今日营业额写入数据库
	if err := s.WriteTodayRevenueByShopID(shopID, revenue); err != nil {
		return 0, errors.New("write today revenue failed")
	}
	return revenue, nil
}

// WriteTodayRevenueByShopID 写入指定店铺的今日营业额
func (s *GormStore) WriteTodayRevenueByShopID(shopID int, revenue float64) error {
	//先看是否已经存在 如果已经存在 则只变更revenue字段
	existing, err := dao.GetShopDailyRevenueByShopIDAndDate(s.db, shopID, time.Now().Truncate(24*time.Hour))
	if err != nil {
		return errors.New("get shop daily revenue failed")
	}
	if existing.ID > 0 {
		if err := dao.UpdateShopDailyRevenueRevenue(s.db, existing, revenue); err != nil {
			return errors.New("update shop daily revenue revenue failed")
		}
		return nil
	}
	//如果不存在 则插入新记录
	if err := dao.CreateShopDailyRevenue(s.db, model.ShopDailyRevenue{
		ShopID:  int64(shopID),
		Date:    time.Now().Truncate(24 * time.Hour),
		Revenue: revenue,
	}); err != nil {
		return errors.New("create shop daily revenue failed")
	}
	return nil
}

// GetAllOrderCountByShopID 获取指定店铺的所有订单数
func (s *GormStore) GetAllOrderCountByShopID(shopID int) ([]bizmodel.OrderCountAll, error) {
	orderCounts, err := dao.GetAllOrderCountByShopID(s.db, shopID)
	if err != nil {
		return nil, errors.New("get all order count failed")
	}
	out := make([]bizmodel.OrderCountAll, 0, len(orderCounts))
	for _, orderCount := range orderCounts {
		out = append(out, bizmodel.OrderCountAll{
			Date:       orderCount.Date,
			OrderCount: orderCount.OrderCount,
		})
	}
	return out, nil
}

// GetAllRevenueByShopID 获取指定店铺的所有营业额记录
func (s *GormStore) GetAllRevenueByShopID(shopID int) ([]bizmodel.RevenueCountAll, error) {
	revenueCounts, err := dao.GetAllOrderCountByShopID(s.db, shopID)
	if err != nil {
		return nil, errors.New("get all revenue count failed")
	}
	out := make([]bizmodel.RevenueCountAll, 0, len(revenueCounts))
	for _, revenueCount := range revenueCounts {
		out = append(out, bizmodel.RevenueCountAll{
			Date:    revenueCount.Date,
			Revenue: revenueCount.Revenue,
		})
	}
	return out, nil
}

// UpdateShopInfo 更新指定店铺的资料
func (s *GormStore) UpdateShopInfo(shopID int, shopName string, deliveryRange float64, deliveryFee float64, businessHours string, shopType int) error {
	if err := dao.UpdateShopInfo(s.db, shopID, shopName, deliveryRange, deliveryFee, businessHours, shopType); err != nil {
		return errors.New("update shop info failed")
	}
	return nil
}

// SaveDish 保存菜品
func (s *GormStore) SaveDish(dishID int, dishName string, price float64, stock int, status int) error {
	if err := dao.SaveDish(s.db, dishID, dishName, price, stock, status); err != nil {
		return errors.New("save dish failed")
	}
	return nil
}

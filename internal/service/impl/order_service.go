package impl

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/dao"
	"Order5003/internal/model"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func (s *GormStore) CreateOrder(order bizmodel.Order) bizmodel.Order {
	e := &model.OrderEntity{}
	if err := dao.CreateOrder(s.db, e); err != nil {
		return order
	}
	order.OrderID = e.OrderID
	//CREATE ORDER ZHUHSHI
	return order
}

func (s *GormStore) GetOrderByID(id int) (bizmodel.Order, error) {
	e, err := dao.GetOrderByID(s.db, id)
	if err != nil {
		return bizmodel.Order{}, errors.New("order not found")
	}
	return bizmodel.Order{
		OrderID:     e.OrderID,
		UserID:      e.UserID,
		ShopID:      e.ShopID,
		TotalAmount: e.TotalAmount,
		Status:      e.Status,
		CreatedAt:   e.CreatedAt,
	}, nil
}

func (s *GormStore) GetAllOrders() []bizmodel.Order { //取当前商家的所有Order 那我应该是查当前ShopId的Order
	list, err := dao.ListOrders(s.db)
	if err != nil {
		return []bizmodel.Order{}
	}
	out := make([]bizmodel.Order, 0, len(list))
	for _, e := range list {
		out = append(out, bizmodel.Order{
			OrderID:     e.OrderID,
			UserID:      e.UserID,
			ShopID:      e.ShopID,
			TotalAmount: e.TotalAmount,
			Status:      e.Status,
			CreatedAt:   e.CreatedAt,
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

func (s *GormStore) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	// 开启GORM事务（直接返回*gorm.DB类型的事务对象）
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败：%v", tx.Error)
	}
	// 执行业务逻辑（传入GORM事务对象tx）
	if err := fn(tx); err != nil {
		// 业务失败：回滚事务
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			return fmt.Errorf("业务执行失败：%v；回滚事务失败：%v", err, rollbackErr)
		}
		return err
	}
	// 业务成功：提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败：%v", err)
	}
	return nil
}

func (s *GormStore) GetDishByID(ctx context.Context, tx *gorm.DB, dishID int) (*bizmodel.Dishes, error) {
	// 1. 调用Dao层查询数据库（传递事务tx）
	dishEntity, err := dao.GetDishByID(tx, dishID)
	if err != nil {
		return nil, fmt.Errorf("service查询菜品失败：%v", err)
	}
	// 2. 菜品不存在：返回 nil, nil
	if dishEntity == nil {
		return nil, nil
	}
	// 3. 模型转换：model.DishEntity → bizmodel.Dishes（隔离数据库模型和业务模型）
	return &bizmodel.Dishes{
		DishID:   dishEntity.DishID,
		ShopID:   dishEntity.ShopID,
		DishName: dishEntity.DishName,
		Price:    dishEntity.Price,
		Stock:    dishEntity.Stock,
		Status:   int(dishEntity.Status),
	}, nil
}

func (s *GormStore) CreateOrderMaster(ctx context.Context, tx *gorm.DB, orderMaster *bizmodel.Order) (int, error) {
	// 1. 业务模型 → 数据库模型 转换（bizmodel.Order → model.OrderEntity）
	orderEntity := &model.OrderEntity{
		UserID:      orderMaster.UserID,
		ShopID:      orderMaster.ShopID,
		TotalAmount: orderMaster.TotalAmount, // shopspring decimal，GORM自动映射数据库decimal
		Status:      orderMaster.Status,      // 默认为"待支付"
		CreatedAt:   time.Now(),              // 手动赋值（或依赖数据库默认值）
		// 其他可选字段（DelivererID、PayTime等）：默认nil，对应数据库NULL
	}

	// 2. 调用 Dao 层插入（传递事务 tx）
	if err := dao.CreateOrder(tx, orderEntity); err != nil {
		return 0, fmt.Errorf("service创建订单主表失败：%v", err)
	}

	// 3. GORM Create 后，orderEntity.OrderID 已被赋值为自增ID，直接返回
	return orderEntity.OrderID, nil
}

func (s *GormStore) CreateOrderDish(ctx context.Context, tx *gorm.DB, detail *bizmodel.OrderDishDetail) error {
	// 1. 业务模型 → 数据库模型 转换（bizmodel.OrderDishDetail → model.OrderDishEntity）
	OrderDishEntity := &model.OrderDishEntity{
		OrderID:   detail.OrderID,   // 关联的订单ID（已在Handler层赋值）
		DishID:    detail.DishID,    // 菜品ID
		DishName:  detail.DishName,  // 冗余菜品名称
		Quantity:  detail.Quantity,  // 菜品数量
		UnitPrice: detail.UnitPrice, // shopspring decimal，映射数据库unit_price
		Subtotal:  detail.Subtotal,  // 单品小计，映射数据库subtotal
	}

	// 2. 调用 Dao 层插入（传递事务 tx）
	if err := dao.CreateOrderDish(tx, OrderDishEntity); err != nil {
		return fmt.Errorf("service创建订单明细失败：order_id=%d, dish_id=%d, err=%v", detail.OrderID, detail.DishID, err)
	}
	return nil
}

package bizmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderDishItem struct {
	DishId   int `json:"dish_id"`
	Quantity int `json:"quantity"`
}

type Order struct {
	OrderID     int             `json:"order_id"`     // 订单ID（自增）
	UserID      int             `json:"user_id"`      // 下单用户ID
	ShopID      int             `json:"shop_id"`      // 所属商家ID
	TotalAmount decimal.Decimal `json:"total_amount"` // 订单总金额
	Status      string          `json:"status"`       // 订单状态
	CreatedAt   time.Time       `json:"created_at"`   // 创建时间
}

type NewOrderRequest struct {
	UserID int             `json:"user_id"`
	Dishes []OrderDishItem `json:"items"`
}

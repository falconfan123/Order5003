package bizmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderStatus int

// 可以在Orders的comment中查询到具体含义
// 订单状态：0=未下单、1=待支付、2=待接单、3=已自我取消、4=待配送、5=配送中、6=已完成、7=已被商家取消、8=备餐中
const (
	OrderStatusPending         OrderStatus = 1
	OrderStatusPreparing       OrderStatus = 8
	OrderStatusReadyForDeliver OrderStatus = 5
	OrderStatusCompleted       OrderStatus = 6
	OrderStatusSelfCancelled   OrderStatus = 3
	OrderStatusShopCancelled   OrderStatus = 7
)

type OrderDishItem struct {
	DishId   int             `json:"dish_id"`
	Quantity int             `json:"quantity"`
	Price    decimal.Decimal `json:"price"`
}

type Order struct {
	OrderID     int             `json:"order_id"`     // 订单ID（自增）
	UserID      int             `json:"user_id"`      // 下单用户ID
	ShopID      int             `json:"shop_id"`      // 所属商家ID
	TotalAmount decimal.Decimal `json:"total_amount"` // 订单总金额
	Status      int             `json:"status"`       // 订单状态
	CreatedAt   time.Time       `json:"created_at"`   // 创建时间
}

type NewOrderRequest struct {
	UserID int             `json:"user_id"`
	ShopID int             `json:"shop_id"`
	Dishes []OrderDishItem `json:"dishes"`
}

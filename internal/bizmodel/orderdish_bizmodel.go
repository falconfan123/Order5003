package bizmodel

import "github.com/shopspring/decimal"

type OrderDishDetail struct {
	OrderID   int             `json:"order_id,omitempty"` // 订单ID（响应时返回，请求时无需传）
	DishID    int             `json:"dish_id"`            // 菜品ID
	DishName  string          `json:"dish_name"`          // 菜品名称（响应时返回，冗余展示）
	Quantity  int             `json:"quantity"`           // 菜品数量
	UnitPrice decimal.Decimal `json:"unit_price"`         // 下单时单价（响应时返回，便于前端核对）
	Subtotal  decimal.Decimal `json:"subtotal"`           // 单品小计（响应时返回，便于前端核对）
}

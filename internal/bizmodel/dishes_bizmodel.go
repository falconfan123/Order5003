package bizmodel

import "github.com/shopspring/decimal"

const (
	DishStatusAvailable = 1
	DishStatusStopped   = 0
)

type Dishes struct {
	DishID      int             `json:"dish_id"`               // 对应前端dish_id（必传）
	ShopID      int             `json:"shop_id"`               // 对应前端shop_id（必传）
	DishName    string          `json:"dish_name"`             // 对应前端dish_name（必传）
	Price       decimal.Decimal `json:"price"`                 // 对应前端price（必传，需自定义序列化）
	Stock       int             `json:"stock,omitempty"`       // 库存（可选，前端忽略不影响）
	Status      int             `json:"status,omitempty"`      // 菜品状态（可选，1=上架/0=下架）
	Description string          `json:"description,omitempty"` // 前端需要的描述（补充）
	Image       string          `json:"image,omitempty"`       // 前端需要的图片URL（补充）
}

package bizmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type Shop struct {
	ShopID        int             `json:"shop_id"`
	ShopName      string          `json:"shop_name"`
	DeliveryRange float64         `json:"delivery_range"`
	DeliveryFee   decimal.Decimal `json:"delivery_fee"`
	BusinessHours string          `json:"business_hours"`
	Status        int             `json:"status"`
	CreatedAt     *time.Time      `json:"created_at"`
	Password      string          `json:"password"`
	Type          int             `json:"type"`
}

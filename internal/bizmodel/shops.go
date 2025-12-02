package bizmodel

import "time"

type Shop struct {
    ShopID        int        `json:"shop_id"`
    ShopName      string     `json:"shop_name"`
    DeliveryRange float64    `json:"delivery_range"`
    DeliveryFee   float64    `json:"delivery_fee"`
    BusinessHours string     `json:"business_hours"`
    Status        int        `json:"status"`
    CreatedAt     *time.Time `json:"created_at"`
    Password      string     `json:"password"`
}


package models

import "time"

type Shop struct {
    ID        int       `json:"id"`
    ShopName  string    `json:"shop_name"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

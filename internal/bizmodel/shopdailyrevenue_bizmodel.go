package bizmodel

import (
	"time"
)

type OrderCountAll struct {
	Date       time.Time `json:"date"`
	OrderCount int       `json:"order_count"`
}

type RevenueCountAll struct {
	Date    time.Time `json:"date"`
	Revenue float64   `json:"revenue"`
}

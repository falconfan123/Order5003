package bizmodel

import "time"

type OrderStatus string

const (
    OrderStatusPending   OrderStatus = "pending"
    OrderStatusPreparing OrderStatus = "preparing"
    OrderStatusReady     OrderStatus = "ready"
    OrderStatusCompleted OrderStatus = "completed"
    OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderItem struct {
    MenuItemID int     `json:"menu_item_id"`
    Quantity   int     `json:"quantity"`
    Price      float64 `json:"price"`
    Name       string  `json:"name"`
}

type Order struct {
    ID        int         `json:"id"`
    TableNumber string    `json:"table_number"`
    Items     []OrderItem `json:"items"`
    Total     float64     `json:"total"`
    Status    OrderStatus `json:"status"`
    CreatedAt time.Time   `json:"created_at"`
    UpdatedAt time.Time   `json:"updated_at"`
}

type NewOrderRequest struct {
    TableNumber string      `json:"table_number"`
    Items       []OrderItem `json:"items"`
}


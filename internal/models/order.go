package models

import "time"

// OrderStatus 表示订单的状态类型
type OrderStatus string

const (
    OrderStatusPending    OrderStatus = "pending"     // 待处理
    OrderStatusPreparing  OrderStatus = "preparing"   // 制作中
    OrderStatusReady      OrderStatus = "ready"       // 已就绪
    OrderStatusCompleted  OrderStatus = "completed"   // 已完成
    OrderStatusCancelled  OrderStatus = "cancelled"   // 已取消
)

// OrderItem 表示订单中的一个菜品项
type OrderItem struct {
    MenuItemID int     `json:"menu_item_id"`
    Quantity   int     `json:"quantity"`
    Price      float64 `json:"price"`
    Name       string  `json:"name"`
}

// Order 表示一个完整的订单
type Order struct {
    ID          int         `json:"id"`
    TableNumber string      `json:"table_number"`
    Items       []OrderItem `json:"items"`
    Total       float64     `json:"total"`
    Status      OrderStatus `json:"status"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

// NewOrderRequest 表示创建新订单的请求数据结构
type NewOrderRequest struct {
    TableNumber string      `json:"table_number"`
    Items       []OrderItem `json:"items"`
}
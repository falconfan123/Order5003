package model

import "time"

type OrderEntity struct {
    OrderID      int        `gorm:"column:order_id;primaryKey;autoIncrement"`
    UserID       int        `gorm:"column:user_id;not null"`
    ShopID       int        `gorm:"column:shop_id;not null"`
    DelivererID  *int       `gorm:"column:deliverer_id"`
    TotalAmount  float64    `gorm:"column:total_amount;not null"`
    Status       string     `gorm:"column:status;type:varchar(20);not null;default:'未下单'"`
    CreatedAt    time.Time  `gorm:"column:created_at"`
    PayTime      *time.Time `gorm:"column:pay_time"`
    RejectReason *string    `gorm:"column:reject_reason;type:varchar(200)"`
    CompletedAt  *time.Time `gorm:"column:completed_at"`
}

func (OrderEntity) TableName() string { return "orders" }

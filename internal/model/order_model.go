package model

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type OrderEntity struct {
	OrderID int `gorm:"column:order_id;primaryKey;autoIncrement"`
	UserID  int `gorm:"column:user_id;not null"`
	ShopID  int `gorm:"column:shop_id;not null"`
	// 关键修改：改为sql.NullInt64，允许NULL，默认NULL
	DelivererID  sql.NullInt64   `gorm:"column:deliverer_id;null;default:NULL"`
	TotalAmount  decimal.Decimal `gorm:"column:total_amount;not null"`
	Status       int             `gorm:"column:status;type:int;not null;default:1"`
	CreatedAt    time.Time       `gorm:"column:created_at;autoCreateTime"`
	PayTime      sql.NullTime    `gorm:"column:pay_time;null;default:NULL"`
	RejectReason string          `gorm:"column:reject_reason;type:varchar(200);null;default:''"`
	CompletedAt  sql.NullTime    `gorm:"column:completed_at;null;default:NULL"`
}

func (OrderEntity) TableName() string { return "orders" }

package model

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

type OrderEntity struct {
	OrderID      int             `gorm:"column:order_id;primaryKey;autoIncrement"`
	UserID       int             `gorm:"column:user_id;not null"`
	ShopID       int             `gorm:"column:shop_id;not null"`
	DelivererID  int             `gorm:"column:deliverer_id;default:0"` // 补充：配送员ID默认0
	TotalAmount  decimal.Decimal `gorm:"column:total_amount;not null"`
	Status       int             `gorm:"column:status;type:int;not null;default:1"` // 修正：默认1（待支付/待接单），而非0
	CreatedAt    time.Time       `gorm:"column:created_at;autoCreateTime"`          // 自动填充创建时间
	PayTime      sql.NullTime    `gorm:"column:pay_time;null;default:NULL"`         // 未支付时为NULL
	RejectReason string          `gorm:"column:reject_reason;type:varchar(200);null;default:''"`
	CompletedAt  sql.NullTime    `gorm:"column:completed_at;null;default:NULL"` // 未完成时为NULL
}

func (OrderEntity) TableName() string { return "orders" }

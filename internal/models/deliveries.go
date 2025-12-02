package models

import "time"

type Deliveries struct {
    DeliveryID     int        `gorm:"column:delivery_id;primaryKey"`
    OrderID        int        `gorm:"column:order_id"`
    DelivererID    int        `gorm:"column:deliverer_id"`
    PickUpTime     *time.Time `gorm:"column:pick_up_time"`
    DeliverTime    *time.Time `gorm:"column:deliver_time"`
    DeliveryStatus int        `gorm:"column:delivery_status"`
}

func (Deliveries) TableName() string { return "deliveries" }

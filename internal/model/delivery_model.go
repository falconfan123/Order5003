package model

import "time"

type DeliveryEntity struct {
    DeliveryID     int       `gorm:"column:delivery_id;primaryKey;autoIncrement" comment:"配送记录唯一ID"`
    OrderID        int       `gorm:"column:order_id;not null" comment:"关联订单ID（关联orders表order_id）"`
    DelivererID    int       `gorm:"column:deliverer_id;not null" comment:"负责配送的配送员ID（关联deliverers表deliverer_id）"`
    PickUpTime     time.Time `gorm:"column:pick_up_time" comment:"取餐时间"`
    DeliverTime    time.Time `gorm:"column:deliver_time" comment:"送达时间"`
    DeliveryStatus int8      `gorm:"column:delivery_status;not null" comment:"配送状态（0=待取餐，1=配送中，2=已送达）"`
}

func (DeliveryEntity) TableName() string { return "deliveries" }

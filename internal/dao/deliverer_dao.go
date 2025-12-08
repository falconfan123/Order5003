package dao

import (
	"Order5003/internal/model"
	"time"

	"gorm.io/gorm"
)

func GetDelivererByName(db *gorm.DB, name string) (*model.DelivererEntity, error) {
	var e model.DelivererEntity
	if err := db.Where("name = ?", name).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func AcceptOrder(db *gorm.DB, deliverID int, orderID int) error {
	//更改deliveries表下的信息
	Deliveries := model.DeliveryEntity{
		DelivererID:    deliverID,
		OrderID:        orderID,
		PickUpTime:     time.Now(),
		DeliverTime:    time.Now(),
		DeliveryStatus: 1,
	}
	if err := db.Save(&Deliveries).Error; err != nil {
		return err
	}
	return nil
}

func GetMyOrder(db *gorm.DB, deliverID int) ([]model.DeliveryEntity, error) {
	var deliveries []model.DeliveryEntity
	if err := db.Where("deliverer_id = ?", deliverID).Find(&deliveries).Error; err != nil {
		return nil, err
	}
	return deliveries, nil
}

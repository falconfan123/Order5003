package dao

import (
	"Order5003/internal/model"

	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB, e *model.OrderEntity) error {
	if err := db.Create(e).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderByID(db *gorm.DB, id int) (*model.OrderEntity, error) {
	var e model.OrderEntity
	if err := db.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func ListOrders(db *gorm.DB) ([]model.OrderEntity, error) {
	var list []model.OrderEntity
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func UpdateOrderStatus(db *gorm.DB, id int, status string) error {
	return db.Model(&model.OrderEntity{}).Where("order_id = ?", id).Update("status", status).Error
}

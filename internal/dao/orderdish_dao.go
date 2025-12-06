package dao

import (
	"Order5003/internal/model"

	"gorm.io/gorm"
)

func CreateOrderDish(db *gorm.DB, e *model.OrderDishEntity) error {
	if err := db.Create(e).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderDishesByOrderID(db *gorm.DB, orderID int) ([]model.OrderDishEntity, error) {
	var list []model.OrderDishEntity
	if err := db.Where("order_id = ?", orderID).Find(&list).Error; err != nil {
		return []model.OrderDishEntity{}, err
	}
	return list, nil
}

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

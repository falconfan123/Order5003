package dao

import (
	"Order5003/internal/model"

	"gorm.io/gorm"
)

func GetDishesByMenuID(db *gorm.DB, menuID int) ([]model.MenuDishesEntity, error) {
	var list []model.MenuDishesEntity
	if err := db.Where("menu_id = ?", menuID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

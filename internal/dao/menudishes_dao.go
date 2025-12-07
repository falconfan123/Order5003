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

func GetDishesIDByMenuID(db *gorm.DB, menuID int) ([]int, error) {
	var list []model.MenuDishesEntity
	if err := db.Where("menu_id = ?", menuID).Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]int, 0, len(list))
	for _, e := range list {
		out = append(out, e.DishID)
	}
	return out, nil
}

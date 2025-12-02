package dao

import (
	"Order5003/internal/model"

	"gorm.io/gorm"
)

func GetMenuByShopID(db *gorm.DB, shopID int) ([]model.MenuEntity, error) {
	var list []model.MenuEntity
	err := db.Where("shop_id = ?", shopID).Find(&list).Error
	if err != nil {
		return []model.MenuEntity{}, err
	}
	return list, nil
}

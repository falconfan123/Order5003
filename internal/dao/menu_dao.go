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

// GetMenuIDByMenuName 根据菜单名称获取菜单ID
func GetMenuIDByMenuName(db *gorm.DB, shopID int, menuName string) (int, error) {
	var menu model.MenuEntity
	err := db.Where("shop_id = ? AND menu_name = ?", shopID, menuName).First(&menu).Error
	if err != nil {
		return 0, err
	}
	return menu.MenuID, nil
}

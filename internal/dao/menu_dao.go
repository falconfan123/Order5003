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

// CreateMenu 创建菜单
func CreateMenu(db *gorm.DB, menu model.MenuEntity) error {
	err := db.Create(&menu).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteMenu 删除菜单
func DeleteMenu(db *gorm.DB, menuID int) error {
	// 开启事务
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	// 步骤1：删除menu_dishes中关联的记录
	if err := tx.Where("menu_id = ?", menuID).Delete(&model.MenuDishesEntity{}).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}
	// 步骤2：删除menu表中的记录
	if err := tx.Delete(&model.MenuEntity{}, menuID).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}
	// 提交事务
	return tx.Commit().Error
}

// UpdateMenu 更新菜单
func UpdateMenu(db *gorm.DB, menuID int, menuName string, status int8) error {
	err := db.Model(&model.MenuEntity{}).
		Where("menu_id = ?", menuID).
		Updates(map[string]interface{}{
			"menu_name": menuName,
			"status":    status,
		}).Error
	if err != nil {
		return err
	}
	return nil
}

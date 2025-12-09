package dao

import (
	"Order5003/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

// AddDishToMenu 添加菜品到菜单
func AddDishToMenu(db *gorm.DB, menuID int, dishID int) error {
	// 步骤1：查询当前菜单下最大sort值（无记录时默认0）
	var maxSort int
	err := db.Model(&model.MenuDishesEntity{}).
		Where("menu_id = ?", menuID).
		Select("IFNULL(MAX(sort), 0)").
		Scan(&maxSort).Error
	if err != nil {
		return err
	}
	// 步骤2：计算新菜品的排序值（最大sort+1，默认排最后）
	newSort := maxSort + 1
	// 步骤3：构建关联记录
	menuDish := &model.MenuDishesEntity{
		MenuID: menuID,
		DishID: dishID,
		Sort:   newSort, // 赋值计算后的排序值
		Status: 1,       // 状态默认显示（符合表默认值）
	}
	// 步骤4：创建记录，处理唯一键冲突（uk_menu_dish）
	// 冲突时更新sort为最新计算的newSort，其他字段保持不变
	return db.Clauses(clause.OnConflict{
		OnConstraint: "uk_menu_dish", // 菜单+菜品唯一键
		DoUpdates: clause.Assignments(map[string]interface{}{
			"sort": newSort, // 替换未定义的sortVal为newSort
		}),
	}).Create(menuDish).Error
}

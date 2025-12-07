package dao

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/model"
	"context"

	"gorm.io/gorm"
)

func ListDishes(db *gorm.DB) ([]model.DishEntity, error) {
	var list []model.DishEntity
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// 只取上架的
func GetDishByDishID(db *gorm.DB, dishID int) (*model.DishEntity, error) {
	var e model.DishEntity
	if err := db.Where("dish_id = ? AND status = 1", dishID).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

// 上架下架都取并返回状态
func GetBothDishByDishID(ctx context.Context, db *gorm.DB, dishID int) (*bizmodel.Dishes, error) {
	var e model.DishEntity
	if err := db.Where("dish_id = ?", dishID).First(&e).Error; err != nil {
		return nil, err
	}
	return &bizmodel.Dishes{
		DishID:   e.DishID,
		DishName: e.DishName,
		Price:    e.Price,
		Status:   int(e.Status),
		Stock:    e.Stock,
		ShopID:   e.ShopID,
	}, nil
}

func CreateDish(db *gorm.DB, e *model.DishEntity) error {
	if err := db.Create(e).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDish(db *gorm.DB, id int, e *model.DishEntity) (*model.DishEntity, error) {
	e.DishID = id
	if err := db.Model(&model.DishEntity{}).Where("dish_id = ?", id).Updates(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func DeleteDish(db *gorm.DB, id int) error {
	return db.Delete(&model.DishEntity{}, id).Error
}

func GetDishesByIDs(db *gorm.DB, dishIDs []int) ([]model.DishEntity, error) {
	var dishes []model.DishEntity
	err := db.Where("dish_id IN (?) AND status = 1", dishIDs).Find(&dishes).Error
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

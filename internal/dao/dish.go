package dao

import (
	"Order5003/internal/model"

	"gorm.io/gorm"
)

func ListDishes(db *gorm.DB) ([]model.DishEntity, error) {
	var list []model.DishEntity
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func GetDishByID(db *gorm.DB, id int) (*model.DishEntity, error) {
	var e model.DishEntity
	if err := db.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func CreateDish(db *gorm.DB, e *model.DishEntity) (*model.DishEntity, error) {
	if err := db.Create(e).Error; err != nil {
		return nil, err
	}
	return e, nil
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

package dao

import "gorm.io/gorm"

func ListDishes(db *gorm.DB) ([]DishEntity, error) {
	var list []DishEntity
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func GetDishByID(db *gorm.DB, id int) (*DishEntity, error) {
	var e DishEntity
	if err := db.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func CreateDish(db *gorm.DB, e *DishEntity) (*DishEntity, error) {
	if err := db.Create(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func UpdateDish(db *gorm.DB, id int, e *DishEntity) (*DishEntity, error) {
	e.DishID = id
	if err := db.Model(&DishEntity{}).Where("id = ?", id).Updates(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func DeleteDish(db *gorm.DB, id int) error {
	return db.Delete(&DishEntity{}, id).Error
}

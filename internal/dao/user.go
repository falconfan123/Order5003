package dao

import (
    "Order5003/internal/model"

    "gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, id int) (*model.UserEntity, error) {
	var e model.UserEntity
	if err := db.Where("user_id = ?", id).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func GetUserByUsername(db *gorm.DB, username string) (*model.UserEntity, error) {
	var e model.UserEntity
	if err := db.Where("user_name = ?", username).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

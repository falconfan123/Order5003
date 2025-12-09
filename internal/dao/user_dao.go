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

func GetUserAddressByUserID(db *gorm.DB, userID int) (*model.UserEntity, error) {
	var e model.UserEntity
	if err := db.Where("user_id = ?", userID).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

// GetUserPhoneByUserID 根据用户ID查询用户电话
func GetUserPhoneByUserID(db *gorm.DB, userID int) (*model.UserEntity, error) {
	var e model.UserEntity
	if err := db.Where("user_id = ?", userID).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

// UpdateUserAddressByUserID 更新用户ID对应的地址
func UpdateUserAddressByUserID(db *gorm.DB, userID int, address string) error {
	if err := db.Model(&model.UserEntity{}).Where("user_id = ?", userID).Update("main_address", address).Error; err != nil {
		return err
	}
	return nil
}

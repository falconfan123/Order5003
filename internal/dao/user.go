package dao

import "gorm.io/gorm"

func GetUserByID(db *gorm.DB, id int) (*UserEntity, error) {
    var e UserEntity
    if err := db.Where("user_id = ?", id).First(&e).Error; err != nil {
        return nil, err
    }
    return &e, nil
}

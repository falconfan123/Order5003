package dao

import "gorm.io/gorm"

func GetUserByUsername(db *gorm.DB, username string) (*UserEntity, error) {
    var e UserEntity
    if err := db.Where("username = ?", username).First(&e).Error; err != nil {
        return nil, err
    }
    return &e, nil
}

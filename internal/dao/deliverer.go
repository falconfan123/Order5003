package dao

import (
    "Order5003/internal/model"
    "gorm.io/gorm"
)

func GetDelivererByName(db *gorm.DB, name string) (*model.DelivererEntity, error) {
    var e model.DelivererEntity
    if err := db.Where("name = ?", name).First(&e).Error; err != nil {
        return nil, err
    }
    return &e, nil
}

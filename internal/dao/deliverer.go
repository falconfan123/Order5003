package dao

import "gorm.io/gorm"

func GetDelivererByName(db *gorm.DB, name string) (*DelivererEntity, error) {
    var e DelivererEntity
    if err := db.Where("name = ?", name).First(&e).Error; err != nil {
        return nil, err
    }
    return &e, nil
}

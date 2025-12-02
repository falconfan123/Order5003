package dao

import (
    "Order5003/internal/model"
    "gorm.io/gorm"
)

func GetShopByName(db *gorm.DB, name string) (*model.ShopEntity, error) {
    var e model.ShopEntity
    if err := db.Where("shop_name = ?", name).First(&e).Error; err != nil {
        return nil, err
    }
    return &e, nil
}

func ListShops(db *gorm.DB) ([]model.ShopEntity, error) {
    var list []model.ShopEntity
    if err := db.Find(&list).Error; err != nil {
        return []model.ShopEntity{}, err
    }
    return list, nil
}

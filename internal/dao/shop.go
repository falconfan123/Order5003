package dao

import "gorm.io/gorm"

func GetShopByName(db *gorm.DB, name string) (*ShopEntity, error) {
    var e ShopEntity
    if err := db.Where("shop_name = ?", name).First(&e).Error; err != nil {
        return nil, err
    }
    return &e, nil
}

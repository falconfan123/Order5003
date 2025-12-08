package dao

import (
	"Order5003/internal/model"
	"time"

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

func GetShopByID(db *gorm.DB, id int) (*model.ShopEntity, error) {
	var e model.ShopEntity
	if err := db.Where("shop_id = ?", id).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

// ListOrdersByShopID 获取指定店铺的所有订单
func ListOrdersByShopID(db *gorm.DB, shopID int) ([]model.OrderEntity, error) {
	var list []model.OrderEntity
	if err := db.Where("shop_id = ?", shopID).Find(&list).Error; err != nil {
		return []model.OrderEntity{}, err
	}
	return list, nil
}

func UpdateShopStatus(db *gorm.DB, e *model.ShopEntity, status int) error {
	return db.Model(&model.ShopEntity{}).Where("shop_id = ?", e.ShopID).Update("status", status).Error
}

// GetTodayOrderCountByShopID 获取指定店铺的今日订单数
func GetTodayOrderCountByShopID(db *gorm.DB, shopID int) (int, error) {
	var count int64
	if err := db.Debug().Model(&model.OrderEntity{}).
		Where("shop_id = ? AND created_at >= ? AND created_at < ?", shopID, time.Now().Truncate(24*time.Hour), time.Now().Truncate(24*time.Hour).Add(24*time.Hour)).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// UpdateShopInfo 更新指定店铺的资料
func UpdateShopInfo(db *gorm.DB, shopID int, shopName string, deliveryRange float64, deliveryFee float64, businessHours string, shopType int) error {
	return db.Model(&model.ShopEntity{}).Where("shop_id = ?", shopID).Updates(map[string]interface{}{
		"shop_name":      shopName,
		"delivery_range": deliveryRange,
		"delivery_fee":   deliveryFee,
		"business_hours": businessHours,
		"type":           shopType,
	}).Error
}

// SaveDish 保存菜品
func SaveDish(db *gorm.DB, dishID int, dishName string, price float64, stock int, status int) error {
	return db.Model(&model.DishEntity{}).Where("dish_id = ?", dishID).Updates(map[string]interface{}{
		"dish_name": dishName,
		"price":     price,
		"stock":     stock,
		"status":    status,
	}).Error
}

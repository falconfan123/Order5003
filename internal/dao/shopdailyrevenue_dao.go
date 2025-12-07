package dao

import (
	"Order5003/internal/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

// GetShopDailyRevenueByShopIDAndDate 根据店铺ID和日期查询店铺的日营业额记录
func GetShopDailyRevenueByShopIDAndDate(db *gorm.DB, shopID int, date time.Time) (model.ShopDailyRevenue, error) {
	var existing model.ShopDailyRevenue
	if err := db.Where("shop_id = ? AND date = ?", shopID, date).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ShopDailyRevenue{}, nil
		}
		return model.ShopDailyRevenue{}, err
	}
	return existing, nil
}

// UpdateShopDailyRevenueCount 更新店铺的日营业额记录
func UpdateShopDailyRevenueCount(db *gorm.DB, existing model.ShopDailyRevenue, count int) error {
	existing.OrderCount = count
	if err := db.Save(&existing).Error; err != nil {
		return err
	}
	return nil
}

// CreateShopDailyRevenue 创建店铺的日营业额记录
func CreateShopDailyRevenue(db *gorm.DB, shopDailyRevenue model.ShopDailyRevenue) error {
	if err := db.Create(&shopDailyRevenue).Error; err != nil {
		return err
	}
	return nil
}

package dao

import (
	"Order5003/internal/logger"
	"Order5003/internal/model"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	db.Debug().Model(&model.ShopDailyRevenue{}).
		Where("shop_id = ? AND date = ?", existing.ShopID, existing.Date).
		Update("order_count", count)

	if err := db.Debug().Model(&model.ShopDailyRevenue{}).
		Where("shop_id = ? AND date = ?", existing.ShopID, existing.Date).
		Update("order_count", count).Error; err != nil {
		logger.Error("更新失败", zap.Error(err))
		return err
	}
	return nil
}

// CreateShopDailyRevenue 创建店铺的日营业额记录
func CreateShopDailyRevenue(db *gorm.DB, shopDailyRevenue model.ShopDailyRevenue) error {
	err := db.Clauses(clause.OnConflict{
		OnConstraint: "uk_shop_date",
		DoUpdates: clause.Assignments(map[string]interface{}{
			"order_count": shopDailyRevenue.OrderCount,
		}),
	}).Create(&shopDailyRevenue).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateShopDailyRevenueRevenue 更新店铺的日营业额记录
func UpdateShopDailyRevenueRevenue(db *gorm.DB, existing model.ShopDailyRevenue, revenue float64) error {
	existing.Revenue = revenue
	if err := db.Save(&existing).Error; err != nil {
		return err
	}
	return nil
}

// GetDailyRevenueByShopID 根据店铺ID查询店铺的日营业额记录
func GetDailyRevenueByShopID(db *gorm.DB, shopID int) (float64, error) {
	var revenue float64
	if err := db.Model(&model.ShopDailyRevenue{}).Where("shop_id = ? AND date = ?", shopID, time.Now().Truncate(24*time.Hour)).Pluck("revenue", &revenue).Error; err != nil {
		return 0, err
	}
	return revenue, nil
}

// GetAllOrderCountByShopID 根据店铺ID查询店铺的所有订单数记录
func GetAllOrderCountByShopID(db *gorm.DB, shopID int) ([]model.ShopDailyRevenue, error) {
	var orderCounts []model.ShopDailyRevenue
	if err := db.Model(&model.ShopDailyRevenue{}).Where("shop_id = ?", shopID).Find(&orderCounts).Error; err != nil {
		return nil, err
	}
	return orderCounts, nil
}

package dao

import (
	"Order5003/internal/logger"
	"Order5003/internal/model"
	"time"

	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB, e *model.OrderEntity) error {
	dr := db.Session(&gorm.Session{DryRun: true}).Create(e)
	if dr.Error == nil && dr.Statement != nil && dr.Statement.SQL.Len() > 0 {
		logger.Info("GORM DryRun SQL", dr.Statement.SQL.String(), "VARS", dr.Statement.Vars)
	}
	if err := db.Create(e).Error; err != nil {
		return err
	}
	return nil
}

func GetOrdersByUserID(db *gorm.DB, userID int) ([]model.OrderEntity, error) {
	var e []model.OrderEntity
	if err := db.Where("user_id = ?", userID).Find(&e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func UpdateOrderStatus(db *gorm.DB, id int, status int) error {
	return db.Model(&model.OrderEntity{}).Where("order_id = ?", id).Update("status", status).Error
}

func ListOrders(db *gorm.DB) ([]model.OrderEntity, error) {
	var list []model.OrderEntity
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// GetTodayFinishOrderByShopID 根据店铺ID查询今日完成的订单数
func GetTodayFinishOrderByShopID(db *gorm.DB, shopID int) ([]model.OrderEntity, error) {
	var list []model.OrderEntity
	if err := db.Model(&model.OrderEntity{}).Where("shop_id = ? AND status = ? AND created_at >= ? AND created_at < ?", shopID, model.OrderStatusCompleted, time.Now().Truncate(24*time.Hour), time.Now().Truncate(24*time.Hour).Add(24*time.Hour)).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

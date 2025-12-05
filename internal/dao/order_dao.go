package dao

import (
	"Order5003/internal/logger"
	"Order5003/internal/model"

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

func GetOrderByUserID(db *gorm.DB, userID int) (*model.OrderEntity, error) {
	var e model.OrderEntity
	if err := db.First(&e, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func UpdateOrderStatus(db *gorm.DB, id int, status string) error {
	return db.Model(&model.OrderEntity{}).Where("order_id = ?", id).Update("status", status).Error
}

func ListOrders(db *gorm.DB) ([]model.OrderEntity, error) {
	var list []model.OrderEntity
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

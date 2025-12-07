package dao

import (
	"Order5003/internal/logger"
	"Order5003/internal/model"
	"runtime"
	"strings"
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
	pc, file, line, ok := runtime.Caller(1)
	callerInfo := "unknown caller"
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		// 简化函数名（去掉包路径）
		funcName = funcName[strings.LastIndex(funcName, "/")+1:]
		callerInfo = funcName + " (" + file + ":" + string(rune(line)) + ")"
	}
	logger.Info("GetShopByID id", id, "caller", callerInfo)

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

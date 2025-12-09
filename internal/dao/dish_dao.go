package dao

import (
	"Order5003/internal/bizmodel"
	"Order5003/internal/model"
	"context"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ListDishes(db *gorm.DB) ([]model.DishEntity, error) {
	var list []model.DishEntity
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// 只取上架的
func GetDishByDishID(db *gorm.DB, dishID int) (*model.DishEntity, error) {
	var e model.DishEntity
	if err := db.Where("dish_id = ? AND status = 1", dishID).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

// 上架下架都取并返回状态
func GetBothDishByDishID(ctx context.Context, db *gorm.DB, dishID int) (*bizmodel.Dishes, error) {
	var e model.DishEntity
	if err := db.Where("dish_id = ?", dishID).First(&e).Error; err != nil {
		return nil, err
	}
	return &bizmodel.Dishes{
		DishID:   e.DishID,
		DishName: e.DishName,
		Price:    e.Price,
		Status:   int(e.Status),
		Stock:    e.Stock,
		ShopID:   e.ShopID,
	}, nil
}

func CreateDish(db *gorm.DB, e *model.DishEntity) error {
	if err := db.Create(e).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDish(db *gorm.DB, id int, e *model.DishEntity) (*model.DishEntity, error) {
	e.DishID = id
	if err := db.Model(&model.DishEntity{}).Where("dish_id = ?", id).Updates(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func DeleteDish(db *gorm.DB, id int) error {
	return db.Delete(&model.DishEntity{}, id).Error
}

func GetDishesByIDs(db *gorm.DB, dishIDs []int) ([]model.DishEntity, error) {
	var dishes []model.DishEntity
	err := db.Where("dish_id IN (?) AND status = 1", dishIDs).Find(&dishes).Error
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

// UpdateDishStatus 更新菜品状态
func UpdateDishStatus(db *gorm.DB, dishID int, status int) error {
	return db.Model(&model.DishEntity{}).Where("dish_id = ?", dishID).Update("status", status).Error
}

// AddDish 添加菜品
func AddDish(db *gorm.DB, shopID int, dishName string, price float64, stock int, status int) (int, error) {
	dish := &model.DishEntity{
		ShopID:   shopID,
		DishName: dishName,
		Price:    decimal.NewFromFloat(price),
		Stock:    stock,
		Status:   int8(status),
	}
	// 2. 执行创建操作，Gorm会自动将生成的dish_id赋值给dish.DishID
	if err := db.Create(dish).Error; err != nil {
		return 0, err // 创建失败，返回0和错误
	}
	// 3. 返回生成的dish_id和nil（无错误）
	return dish.DishID, nil
}

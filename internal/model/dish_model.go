package model

import "github.com/shopspring/decimal"

type DishEntity struct {
	DishID   int             `gorm:"column:dish_id;primaryKey;autoIncrement" comment:"菜品唯一ID"`
	ShopID   int             `gorm:"column:shop_id;not null" comment:"所属商家ID（关联shops表shop_id）"`
	DishName string          `gorm:"column:dish_name;not null" comment:"菜品名称"`
	Price    decimal.Decimal `gorm:"column:price;type:decimal(10,2);not null" comment:"菜品单价"`
	Stock    int             `gorm:"column:stock;not null;default:0" comment:"库存数量"`
	Status   int8            `gorm:"column:status;not null;default:1" comment:"菜品状态（1=上架，0=下架）"`
}

func (DishEntity) TableName() string { return "dishes" }

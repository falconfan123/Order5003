package models

type Dishes struct {
	DishID   int     `gorm:"column:dish_id;primaryKey"`
	ShopID   int     `gorm:"column:shop_id"`
	DishName string  `gorm:"column:dish_name"`
	Price    float64 `gorm:"column:price"`
	Stock    int     `gorm:"column:stock"`
	Status   int     `gorm:"column:status"`
}

func (Dishes) TableName() string { return "dishes" }

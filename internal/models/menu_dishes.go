package models

type MenuDishes struct {
	ID     int `gorm:"column:id;primaryKey"`
	MenuID int `gorm:"column:menu_id"`
	DishID int `gorm:"column:dish_id"`
	Sort   int `gorm:"column:sort"`
	Status int `gorm:"column:status"`
}

func (MenuDishes) TableName() string { return "menu_dishes" }

package model

type MenuDishesEntity struct {
    ID     int  `gorm:"column:id;primaryKey;autoIncrement" comment:"关联记录唯一ID"`
    MenuID int  `gorm:"column:menu_id;not null" comment:"关联菜单ID（关联menu表）"`
    DishID int  `gorm:"column:dish_id;not null" comment:"关联菜品ID（关联dishes表）"`
    Sort   int  `gorm:"column:sort;not null;default:0" comment:"菜品在菜单中的排序（数字越小越靠前）"`
    Status int8 `gorm:"column:status;not null;default:1" comment:"菜品在菜单中的状态（1=显示，0=隐藏）"`
}

func (MenuDishesEntity) TableName() string { return "menu_dishes" }

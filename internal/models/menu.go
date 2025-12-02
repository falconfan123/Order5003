package models

import "time"

type Menu struct {
    MenuID     int       `gorm:"column:menu_id;primaryKey"`
    ShopID     int       `gorm:"column:shop_id"`
    MenuName   string    `gorm:"column:menu_name"`
    Status     int       `gorm:"column:status"`
    CreateTime time.Time `gorm:"column:create_time"`
    UpdateTime time.Time `gorm:"column:update_time"`
}

func (Menu) TableName() string { return "menu" }

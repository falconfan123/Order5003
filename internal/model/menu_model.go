package model

import "time"

type MenuEntity struct {
    MenuID     int       `gorm:"column:menu_id;primaryKey;autoIncrement" comment:"菜单唯一ID"`
    ShopID     int       `gorm:"column:shop_id;not null" comment:"所属商家ID（关联shops表）"`
    MenuName   string    `gorm:"column:menu_name;not null" comment:"菜单名称（如：早餐菜单、招牌菜菜单）"`
    Status     int8      `gorm:"column:status;not null;default:1" comment:"菜单状态（1=启用，0=停用）"`
    CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" comment:"创建时间"`
    UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" comment:"更新时间"`
}

func (MenuEntity) TableName() string { return "menu" }

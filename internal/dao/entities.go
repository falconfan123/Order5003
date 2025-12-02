package dao

import (
	"time"
)

// ShopEntity 对应 shops 表（完全对齐数据库字段）
type ShopEntity struct {
	ShopID        int       `gorm:"column:shop_id;primaryKey;autoIncrement" comment:"商家唯一ID"`
	ShopName      string    `gorm:"column:shop_name;not null" comment:"商家名称"`
	DeliveryRange float64   `gorm:"column:delivery_range;not null;default:5.00" comment:"配送范围（单位：km）"`
	DeliveryFee   float64   `gorm:"column:delivery_fee;not null;default:0.00" comment:"配送费"`
	BusinessHours string    `gorm:"column:business_hours;not null" comment:"商家时间（如：09:00-22:00）"`
	Status        int8      `gorm:"column:status;not null;default:1" comment:"商家状态（1=营业，0=休息）"`
	CreatedAt     time.Time `gorm:"column:created_at" comment:"创建时间"`
	Password      string    `gorm:"column:password" comment:"商家明文密码"`
}

func (ShopEntity) TableName() string { return "shops" }

// DelivererEntity 对应 deliverers 表（完全对齐数据库字段）
type DelivererEntity struct {
	DelivererID     int    `gorm:"column:deliverer_id;primaryKey;autoIncrement" comment:"配送员唯一ID"`
	Name            string `gorm:"column:name;not null" comment:"配送员姓名"`
	Phone           string `gorm:"column:phone;not null" comment:"配送员手机号"`
	Status          int8   `gorm:"column:status;not null;default:0" comment:"配送员状态（0=离线，1=在线）"`
	ResponsibleArea string `gorm:"column:responsible_area" comment:"负责配送区域（如：朝阳区建国路）"`
	Password        string `gorm:"column:password" comment:"配送员密码（明文自用）"`
}

func (DelivererEntity) TableName() string { return "deliverers" }

// UserEntity 对应 users 表（修正字段映射错误）
type UserEntity struct {
	UserID      int       `gorm:"column:user_id;primaryKey;autoIncrement" comment:"用户唯一ID"`
	Phone       string    `gorm:"column:phone;not null;uniqueIndex" comment:"手机号"`
	UserName    string    `gorm:"column:user_name" comment:"用户名"` // 修正：Username→UserName（驼峰对齐字段）
	MainAddress string    `gorm:"column:main_address" comment:"常用收货地址"`
	CreatedAt   time.Time `gorm:"column:created_at" comment:"创建时间"` // 修正：column=createdat→created_at（对齐数据库）
	Password    string    `gorm:"column:password" comment:"密码"`
}

func (UserEntity) TableName() string { return "users" }

// DishEntity 对应 dishes 表（完全重构，对齐数据库字段）
type DishEntity struct {
	DishID   int     `gorm:"column:dish_id;primaryKey;autoIncrement" comment:"菜品唯一ID"`
	ShopID   int     `gorm:"column:shop_id;not null" comment:"所属商家ID（关联shops表shop_id）"`
	DishName string  `gorm:"column:dish_name;not null" comment:"菜品名称"`
	Price    float64 `gorm:"column:price;not null" comment:"菜品单价"`
	Stock    int     `gorm:"column:stock;not null;default:0" comment:"库存数量"`
	Status   int8    `gorm:"column:status;not null;default:1" comment:"菜品状态（1=上架，0=下架）"`
}

func (DishEntity) TableName() string { return "dishes" }

// OrderEntity 对应 orders 表（完全重构，对齐数据库字段）
type OrderEntity struct {
	OrderID     int       `gorm:"column:order_id;primaryKey;autoIncrement" comment:"订单唯一ID"`
	UserID      int       `gorm:"column:user_id;not null" comment:"下单用户ID（关联users表user_id）"`
	ShopID      int       `gorm:"column:shop_id;not null" comment:"订单所属商家ID（关联shops表shop_id）"`
	TotalAmount float64   `gorm:"column:total_amount;not null" comment:"订单总金额"`
	Status      int8      `gorm:"column:status;not null" comment:"订单状态（0=待支付，1=待接单，2=待配送，3=配送中，4=已完成，5=已取消）"`
	CreatedAt   time.Time `gorm:"column:created_at" comment:"创建时间"`
	PayTime     time.Time `gorm:"column:pay_time" comment:"支付时间（待支付状态时为NULL）"`
}

func (OrderEntity) TableName() string { return "orders" }

// DeliveryEntity 对应 deliveries 表（修正字段类型，补充约束）
type DeliveryEntity struct {
	DeliveryID     int       `gorm:"column:delivery_id;primaryKey;autoIncrement" comment:"配送记录唯一ID"`
	OrderID        int       `gorm:"column:order_id;not null" comment:"关联订单ID（关联orders表order_id）"`
	DelivererID    int       `gorm:"column:deliverer_id;not null" comment:"负责配送的配送员ID（关联deliverers表deliverer_id）"`
	PickUpTime     time.Time `gorm:"column:pick_up_time" comment:"取餐时间"`
	DeliverTime    time.Time `gorm:"column:deliver_time" comment:"送达时间"`
	DeliveryStatus int8      `gorm:"column:delivery_status;not null" comment:"配送状态（0=待取餐，1=配送中，2=已送达）"`
}

func (DeliveryEntity) TableName() string { return "deliveries" }

// MenuEntity 对应 menu 表（补充缺失的Entity）
type MenuEntity struct {
	MenuID     int       `gorm:"column:menu_id;primaryKey;autoIncrement" comment:"菜单唯一ID"`
	ShopID     int       `gorm:"column:shop_id;not null" comment:"所属商家ID（关联shops表）"`
	MenuName   string    `gorm:"column:menu_name;not null" comment:"菜单名称（如：早餐菜单、招牌菜菜单）"`
	Status     int8      `gorm:"column:status;not null;default:1" comment:"菜单状态（1=启用，0=停用）"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" comment:"创建时间"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" comment:"更新时间"`
}

func (MenuEntity) TableName() string { return "menu" }

// MenuDishesEntity 对应 menu_dishes 表（补充缺失的Entity）
type MenuDishesEntity struct {
	ID     int  `gorm:"column:id;primaryKey;autoIncrement" comment:"关联记录唯一ID"`
	MenuID int  `gorm:"column:menu_id;not null" comment:"关联菜单ID（关联menu表）"`
	DishID int  `gorm:"column:dish_id;not null" comment:"关联菜品ID（关联dishes表）"`
	Sort   int  `gorm:"column:sort;not null;default:0" comment:"菜品在菜单中的排序（数字越小越靠前）"`
	Status int8 `gorm:"column:status;not null;default:1" comment:"菜品在菜单中的状态（1=显示，0=隐藏）"`
}

func (MenuDishesEntity) TableName() string { return "menu_dishes" }

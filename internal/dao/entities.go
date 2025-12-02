package dao

import "time"

type ShopEntity struct {
	ShopID   int    `gorm:"column:shop_id;primaryKey"`
	ShopName string `gorm:"column:shop_name"
`
	Password   string    `gorm:"column:password"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (ShopEntity) TableName() string { return "shops" }

type DelivererEntity struct {
	DelivererID     int    `gorm:"column:deliverer_id;primaryKey"`
	Name            string `gorm:"column:name"`
	Phone           string `gorm:"column:phone"`
	Status          int    `gorm:"column:status"`
	ResponsibleArea string `gorm:"column:responsible_area"`
	Password        string `gorm:"column:password"`
}

func (DelivererEntity) TableName() string { return "deliverers" }

type UserEntity struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Role      string    `gorm:"column:role"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserEntity) TableName() string { return "users" }

type DishEntity struct {
	ID          int       `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Price       float64   `gorm:"column:price"`
	Category    string    `gorm:"column:category"`
	IsAvailable int       `gorm:"column:is_available"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (DishEntity) TableName() string { return "dishes" }

type OrderEntity struct {
	ID          int       `gorm:"column:id;primaryKey"`
	TableNumber string    `gorm:"column:table_number"`
	ItemsJSON   string    `gorm:"column:items_json"`
	Total       float64   `gorm:"column:total"`
	Status      string    `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (OrderEntity) TableName() string { return "orders" }

type DeliveryEntity struct {
	DeliveryID     int       `gorm:"column:delivery_id;primaryKey"`
	OrderID        int       `gorm:"column:order_id"`
	DelivererID    int       `gorm:"column:deliverer_id"`
	PickUpTime     time.Time `gorm:"column:pick_up_time"`
	DeliverTime    time.Time `gorm:"column:deliver_time"`
	DeliveryStatus int       `gorm:"column:delivery_status"`
}

func (DeliveryEntity) TableName() string { return "deliveries" }

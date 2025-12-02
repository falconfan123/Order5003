package impl

import (
    "Order5003/internal/service"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type GormStore struct {
    db *gorm.DB
}

var _ service.MenuService = (*GormStore)(nil)
var _ service.OrderService = (*GormStore)(nil)
var _ service.ShopService = (*GormStore)(nil)
var _ service.UserService = (*GormStore)(nil)
var _ service.DelivererService = (*GormStore)(nil)

func NewGormStore(dsn string) (*GormStore, error) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return &GormStore{db: db}, nil
}

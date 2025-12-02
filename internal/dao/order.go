package dao

import "gorm.io/gorm"

func CreateOrder(db *gorm.DB, e *OrderEntity) (*OrderEntity, error) {
	if err := db.Create(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func GetOrderByID(db *gorm.DB, id int) (*OrderEntity, error) {
	var e OrderEntity
	if err := db.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func ListOrders(db *gorm.DB) ([]OrderEntity, error) {
	var list []OrderEntity
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func UpdateOrderStatus(db *gorm.DB, id int, status string) error {
	return db.Model(&OrderEntity{}).Where("id = ?", id).Update("status", status).Error
}

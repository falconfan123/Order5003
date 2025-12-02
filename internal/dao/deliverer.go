package dao

import (
	"Order5003/internal/models"

	"gorm.io/gorm"
)

func GetDelivererByName(db *gorm.DB, name string) (*models.Deliverers, error) {
	var e models.Deliverers
	if err := db.Where("name = ?", name).First(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

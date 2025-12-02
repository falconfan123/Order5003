package models

type Deliverers struct {
	DelivererID     int    `gorm:"column:deliverer_id;primaryKey"`
	Name            string `gorm:"column:name"`
	Phone           string `gorm:"column:phone"`
	Status          int    `gorm:"column:status"`
	ResponsibleArea string `gorm:"column:responsible_area"`
	Password        string `gorm:"column:password"`
}

func (Deliverers) TableName() string { return "deliverers" }

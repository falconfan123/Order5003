package bizmodel

import "time"

type Menu struct {
	MenuID     int
	ShopID     int
	MenuName   string
	Status     int
	CreateTime time.Time
	UpdateTime time.Time
}

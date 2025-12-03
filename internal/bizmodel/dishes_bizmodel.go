package bizmodel

import "github.com/shopspring/decimal"

type Dishes struct {
	DishID   int
	ShopID   int
	DishName string
	Price    decimal.Decimal
	Stock    int
	Status   int
}

package model

import "github.com/shopspring/decimal"

type OrderDishEntity struct {
	ID        int             `gorm:"column:id;type:int;primaryKey;autoIncrement;comment:'关联记录唯一ID'"`                       // 自增主键
	OrderID   int             `gorm:"column:order_id;type:int;not null;comment:'关联订单ID（关联orders表order_id）'"`                // 关联订单ID
	DishID    int             `gorm:"column:dish_id;type:int;not null;comment:'关联菜品ID（关联dishes表dish_id）'"`                  // 关联菜品ID
	DishName  string          `gorm:"column:dish_name;type:varchar(50);not null;comment:'下单时的菜品名称（冗余存储）'"`                  // 冗余菜品名称
	Quantity  int             `gorm:"column:quantity;type:int;not null;comment:'菜品数量（至少1份）'"`                               // 菜品数量
    UnitPrice decimal.Decimal `gorm:"column:unit_price;type:decimal(10,2);not null;comment:'下单时的菜品单价（冗余存储）'"`               // 冗余单价
    Subtotal  decimal.Decimal `gorm:"column:subtotal;type:decimal(10,2);not null;comment:'该菜品小计金额（quantity * unit_price）'"` // 单品小计
}

// TableName 显式指定数据库表名（与建表语句完全一致）
func (OrderDishEntity) TableName() string {
	return "order_dishes"
}

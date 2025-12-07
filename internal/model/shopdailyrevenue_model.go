package model

import "time"

type ShopDailyRevenue struct {
	ID         int64     `gorm:"primaryKey;autoIncrement;column:id;comment:自增主键"`
	ShopID     int64     `gorm:"column:shop_id;not null;comment:店铺ID，关联shops表的主键"`
	Date       time.Time `gorm:"column:date;not null;type:date;comment:统计日期（如2025-12-07）"`
	Revenue    float64   `gorm:"column:revenue;default:0.00;type:decimal(10,2);comment:当日营业额"` // 移除 not null
	OrderCount int       `gorm:"column:order_count;default:0;comment:当日订单数"`                   // 移除 not null
}

// TableName 指定模型对应的数据库表名
func (s *ShopDailyRevenue) TableName() string {
	return "shop_daily_revenue"
}

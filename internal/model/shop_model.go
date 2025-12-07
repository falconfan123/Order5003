package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ShopEntity struct {
	ShopID        int             `gorm:"column:shop_id;primaryKey;autoIncrement" comment:"商家唯一ID"`
	ShopName      string          `gorm:"column:shop_name;not null" comment:"商家名称"`
	DeliveryRange decimal.Decimal `gorm:"column:delivery_range;not null;default:5.00" comment:"配送范围（单位：km）"`
	DeliveryFee   decimal.Decimal `gorm:"column:delivery_fee;not null;default:0.00" comment:"配送费"`
	BusinessHours string          `gorm:"column:business_hours;not null" comment:"商家时间（如：09:00-22:00）"`
	Status        int8            `gorm:"column:status;not null;default:1" comment:"商家状态（1=营业，0=休息）"`
	CreatedAt     time.Time       `gorm:"column:created_at" comment:"创建时间"`
	Password      string          `gorm:"column:password" comment:"商家明文密码"`
	Type          int8            `gorm:"column:type;not null;default:0" comment:"商家分类类型：0=全部，1=甜品奶茶，2=炸鸡汉堡，3=美味中餐，4=生活百货"`
}

func (ShopEntity) TableName() string { return "shops" }

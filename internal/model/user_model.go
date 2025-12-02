package model

import "time"

type UserEntity struct {
    UserID      int       `gorm:"column:user_id;primaryKey;autoIncrement" comment:"用户唯一ID"`
    Phone       string    `gorm:"column:phone;not null;uniqueIndex" comment:"手机号"`
    UserName    string    `gorm:"column:user_name" comment:"用户名"`
    MainAddress string    `gorm:"column:main_address" comment:"常用收货地址"`
    CreatedAt   time.Time `gorm:"column:created_at" comment:"创建时间"`
    Password    string    `gorm:"column:password" comment:"密码"`
}

func (UserEntity) TableName() string { return "users" }

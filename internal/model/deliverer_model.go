package model

type DelivererEntity struct {
    DelivererID     int    `gorm:"column:deliverer_id;primaryKey;autoIncrement" comment:"配送员唯一ID"`
    Name            string `gorm:"column:name;not null" comment:"配送员姓名"`
    Phone           string `gorm:"column:phone;not null" comment:"配送员手机号"`
    Status          int8   `gorm:"column:status;not null;default:0" comment:"配送员状态（0=离线，1=在线）"`
    ResponsibleArea string `gorm:"column:responsible_area" comment:"负责配送区域（如：朝阳区建国路）"`
    Password        string `gorm:"column:password" comment:"配送员密码（明文自用）"`
}

func (DelivererEntity) TableName() string { return "deliverers" }

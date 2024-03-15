package models

import "time"

type Macheine struct {
	IP         string    `form:"ip" json:"ip" gorm:"ip"`
	Password   string    `form:"password" json:"password" gorm:"password"`
	ModifyTime time.Time `form:"modifyTime" json:"modifyTime" gorm:"modifyTime"`
	AddTime    time.Time `form:"addTime" json:"addTime" gorm:"addTime"`
	IsDeleted  bool      `form:"isDeleted" json:"isDeleted" gorm:"isDeleted"`
}

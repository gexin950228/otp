package models

import "time"

type Machine struct {
	IP         string    `form:"ip" json:"ip" gorm:"ip"`
	Password   string    `form:"password" json:"password" gorm:"password"`
	ModifyTime time.Time `form:"modifyTime" json:"modifyTime" gorm:"modifyTime"`
	AddTime    time.Time `form:"addTime" json:"addTime" gorm:"addTime"`
	Maintainer string    `form:"maintainer" json:"maintainer"`
	Department string    `form:"department" json:"department" gorm:"department"`
	DirectLink bool      `form:"directLink" json:"directLink" gorm:"directLink"`
	IsDeleted  bool      `form:"isDeleted" json:"isDeleted" gorm:"isDeleted"`
	VPNInfo    VPNInfo
}

type VPNInfo struct {
	vpnAddress  string
	User        string `form:""`
	Password    string `form:"password" json:"password" gorm:""`
	ConnectTool string `form:"connectTool" json:"connectTool" gorm:"connectTool"`
}

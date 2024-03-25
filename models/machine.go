package models

type Machine struct {
	IP         string `form:"ip" json:"ip" gorm:"ip"`
	Password   string `form:"password" json:"password" gorm:"password"`
	ModifyTime string `form:"modifyTime" json:"modifyTime" gorm:"modifyTime"`
	AddTime    string `form:"addTime" json:"addTime" gorm:"addTime"`
	Maintainer string `form:"maintainer" json:"maintainer" gorm:"modifyUser"`
	ModifyUser string `form:"modifyUser" json:"modifyUser" gorm:"modifyUser"`
	Department string `form:"department" json:"department" gorm:"department"`
	DirectLink bool   `form:"directLink" json:"directLink" gorm:"directLink"`
	IsDeleted  bool   `form:"isDeleted" json:"isDeleted" gorm:"isDeleted"`
	VPN        bool   `form:"vpn" json:"vpn" gorm:"vpn"`
	VPNInfo    VPNInfo
}

type VPNInfo struct {
	VPNAddress     string `form:"vpnAddress" json:"vpnAddress" gorm:"column_name:vonAddress"`
	VPNUser        string `form:"vpnUser" json:"vpnUser" gorm:"column_name:vpnUser"`
	VPNPassword    string `form:"password" json:"password" gorm:"column:vpnPassword"`
	VPNConnectTool string `form:"connectTool" json:"connectTool" gorm:"column_name:connectTool"`
}

package models

type UserInfo struct {
	Id       int    `form:"id" json:"id" gorm:"id"`
	UserName string `form:"username" json:"username" gorm:"username"`
	Password string `form:"password" json:"password" gorm:"password"`
	Email    string `form:"email" json:"email" gorm:"email"`
}

type UserLogin struct {
	UserName    string `form:"username" json:"userName" gorm:"username"`
	Password    string `form:"password" json:"password" gorm:"password"`
	VerifyCode  string `form:"verifyCode" json:"verifyCode" gorm:"verifyCode"`
	RedirectUri string `form:"redirectUri" json:"redirectUri" gorm:"redirectUri"`
}

package models

type UserInfo struct {
	UserName string `form:"username" json:"username" gorm:"username"`
	Password string `form:"password" json:"password" gorm:"password"`
	Email    string `form:"email" json:"email" gorm:"email"`
}

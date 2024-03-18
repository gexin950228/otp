package user

import "github.com/gin-gonic/gin"

func Router(user *gin.RouterGroup) {
	user.GET("/to_login", ToLogin)
	user.POST("/login", Login)
	user.POST("/preLogin", SendLoginVerifyCode)

}

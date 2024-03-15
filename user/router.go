package user

import "github.com/gin-gonic/gin"

func Router(login *gin.RouterGroup) {
	login.GET("/login", Login)
	login.POST("/preLogin", PreLogin)

}

package user

import "github.com/gin-gonic/gin"

func Router(user *gin.RouterGroup) {
	user.GET("/login", Login)
	user.POST("/preLogin", PreLogin)

}

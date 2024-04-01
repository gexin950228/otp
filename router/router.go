package router

import (
	"github.com/gin-gonic/gin"
	"otp/controller"
	"otp/middleware"
	"otp/user"
)

func Router(router *gin.Engine) {
	router.Use(middleware.ParseToken())
	crudRouter := router.Group("/response")
	userRouter := router.Group("/user")
	controller.Router(crudRouter)
	user.Router(userRouter)
}

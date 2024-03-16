package router

import (
	"github.com/gin-gonic/gin"
	"otp/controller"
	"otp/user"
)

func Router(router *gin.Engine) {
	crudRouter := router.Group("/response")
	userRouter := router.Group("/user")
	controller.Router(crudRouter)
	user.Router(userRouter)
}

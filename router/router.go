package router

import (
	"github.com/gin-gonic/gin"
	"otp/responseHandler"
	"otp/user"
)

func Router(router *gin.Engine) {
	crudRouter := router.Group("/response")
	userRouter := router.Group("/user")
	responseHandler.Router(crudRouter)
	user.Router(userRouter)
}

package router

import (
	"github.com/gin-gonic/gin"
	"otp/responseHandler"
)

func Router(router *gin.Engine) {
	response := router.Group("/responseHandler")
	responseHandler.Router(response)
}

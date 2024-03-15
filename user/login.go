package user

import (
	"github.com/gin-gonic/gin"
	"otp/models"
)

func Login(ctx *gin.Context) {
	var userInfo models.UserInfo
	ctx.ShouldBind(&userInfo)
}

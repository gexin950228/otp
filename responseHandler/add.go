package responseHandler

import (
	"github.com/gin-gonic/gin"
	"otp/models"
)

func AddMachine(ctx *gin.Context) {
	var machine models.Machine
	machine.IP = ctx.PostForm("ip")
}

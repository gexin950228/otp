package responseHandler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchRequest(ctx *gin.Context) {
	session := sessions.Default(ctx)
	infoSetTime := session.Get("setTime")
	fmt.Println(infoSetTime)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": infoSetTime,
	})
}

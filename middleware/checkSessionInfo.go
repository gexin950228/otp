package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckSessionInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		//now := time.Now()

	}
}

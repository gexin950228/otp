package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otp/logSource"
)

func Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Success!")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Hello)
	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		logSource.Log.Panic(err.Error())
		return
	}
}

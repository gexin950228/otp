package controller

import "github.com/gin-gonic/gin"

func Router(response *gin.RouterGroup) {
	response.GET("/show/*id", Show)
	response.POST("/add", AddMachine)
	response.POST("/modify", ModifyMachine)
	response.GET("/search/:username", Search)
	response.POST("/delete", DeleteMachine)
}

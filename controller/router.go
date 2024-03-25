package controller

import "github.com/gin-gonic/gin"

func Router(response *gin.RouterGroup) {
	response.GET("/show/*id", Show)
	response.POST("/add", AddMachine)
	response.POST("/modify", ModifyMachine)
	response.POST("/search", Search)
	response.POST("/delete", DeleteMachine)
}

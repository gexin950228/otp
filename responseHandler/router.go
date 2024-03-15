package responseHandler

import "github.com/gin-gonic/gin"

func Router(response *gin.RouterGroup) {
	response.GET("/search", SearchRequest)
}

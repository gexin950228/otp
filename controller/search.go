package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Show(ctx *gin.Context) {
	user := strings.TrimPrefix(ctx.Param("user"), "/")
	data := map[string]interface{}{
		"user": user,
	}
	ctx.HTML(http.StatusOK, "response/show.html", data)
}

func Search(ctx *gin.Context) {
	user := ctx.Param("user")
	fmt.Println(user)
	username := strings.TrimPrefix(user, "/")
	fmt.Printf("username: %s\n", user)
	ctx.HTML(http.StatusOK, "response/show.html", username)
}

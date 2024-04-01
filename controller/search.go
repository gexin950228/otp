package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Show(ctx *gin.Context) {
	Authorization := ctx.Request.Header.Get("Authorization")
	if Authorization == "" {
		ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/user/to_login?uri=/response/show")
	} else {
		ctx.HTML(http.StatusOK, "response/show.html", nil)
	}

}

func Search(ctx *gin.Context) {
	username := ctx.Param("username")
	fmt.Println(username)
	ctx.HTML(http.StatusOK, "response/show.html", username)
}

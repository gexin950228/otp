package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"otp/dataSource"
	"otp/models"
	"otp/sessionInit"
)

func Show(ctx *gin.Context) {
	var machines []models.Machine
	ctx.HTML(http.StatusOK, "template/show.html", machines)
}

func Search(ctx *gin.Context) {
	user := ctx.PostForm("user")
	store := sessionInit.InitSession()
	session := sessions.Default(ctx)
	sessions.Sessions("loginSession", store)

	loginStatus := session.Get(user)
	if loginStatus != "Success" {
		ctx.Redirect(http.StatusUnauthorized, "/user/login")
	}
	ip := ctx.PostForm("ip")
	department := ctx.PostForm("department")
	machineInfo := dataSource.SearchMachine(department, ip)
	if len(machineInfo) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "没找到对应机器，请核对输入是否正确",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
	})
}

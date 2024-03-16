package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"otp/dataSource"
	"otp/models"
	"otp/sessionInit"
)

func DeleteMachine(ctx *gin.Context) {
	var machineDelete models.Machine
	errBind := ctx.ShouldBindJSON(&machineDelete)
	if errBind != nil {
		logrus.Error(fmt.Sprintf("删除机器时参数绑定错误： %s\n", errBind.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 2,
			"msg":  "参数绑定错误",
		})
	}
	user := machineDelete.ModifyUser
	store := sessionInit.InitSession()
	session := sessions.Default(ctx)
	sessions.Sessions("loginSession", store)

	loginStatus := session.Get(user)
	if loginStatus != "Success" {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
	}
	Db := dataSource.Db
	Db.Model(&machineDelete).Update("isDeleted", true)
}

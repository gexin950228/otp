package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"otp/models"
	"otp/sessionInit"
	"time"
)

func ModifyMachine(ctx *gin.Context) {
	var machineModify models.Machine
	store := sessionInit.InitSession()
	session := sessions.Default(ctx)
	sessions.Sessions("loginSession", store)
	errBindModify := ctx.ShouldBind(&machineModify)
	loginStatus := session.Get(machineModify.ModifyUser)
	if loginStatus != "Success" {
		ctx.Redirect(http.StatusUnauthorized, "/user/login")
	}
	machineModify.ModifyTime = time.Now().Format("2006/01/02 15:04:05")
	if errBindModify != nil {
		logrus.Error(fmt.Sprintf("修改机器信息是前端数据绑定出错：%s", errBindModify.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 2,
			"msg":  "提交参数错误，数据绑定出错",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "修改机器成功",
		})
	}

}

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
	"otp/utils"
	"time"
)

var Db = dataSource.InitDb()

func AddMachine(ctx *gin.Context) {
	var machineAdd models.Machine
	var verifyOK bool
	errBindMachine := ctx.ShouldBind(&machineAdd)
	if errBindMachine != nil {
		logrus.Error(fmt.Sprintf("机器信息绑定错误: %s\n", errBindMachine))
		return
	}
	user := machineAdd.ModifyUser
	store := sessionInit.InitSession()
	session := sessions.Default(ctx)
	sessions.Sessions("loginSession", store)

	loginStatus := session.Get(user)
	if loginStatus != "Success" {
		ctx.Redirect(http.StatusUnauthorized, "/user/login")
	}
	verifyOK = utils.IsIP(machineAdd)
	if verifyOK {
		machineAdd.AddTime = time.Now().Format("2006/01/02 15:04:05")
		machineAdd.ModifyTime = time.Now().Format("2006/01/02 15:04:05")
		Db.Create(&machineAdd)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "机器添加成功",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 2,
			"msg":  "提交的机器信息格式不正确，新增机器失败",
		})
	}

}

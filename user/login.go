package user

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"otp/SendMail"
	"otp/dataSource"
	"otp/logSource"
	"otp/models"
	"strconv"
	"time"
)

func PreLogin(ctx *gin.Context) {
	var user models.UserInfo
	err := ctx.ShouldBind(&user)
	if err != nil {
		logSource.Log.Error("发送邮箱验证码时登录信息绑定错误")
		return
	}
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	randomNumber := rng.Intn(1000000)
	body := fmt.Sprintf("您此次登陆的验证码是：%s\n", strconv.Itoa(randomNumber))
	to := user.Email
	Subject := "查询机器密码的登录验证码"
	userId := user.Id
	SendMail.SendEmail(userId, Subject, body, to)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Gexin..950228",
		DB:       0,
	})
	key := strconv.FormatInt(int64(user.Id), 64) + "-" + "loginVerifyCode"
	value := strconv.FormatInt(int64(randomNumber), 64)
	errSaveVerifyCodeToRedis := rdb.Set(context.Background(), key, value, 120).Err()
	if errSaveVerifyCodeToRedis != nil {
		logSource.Log.Error(errSaveVerifyCodeToRedis.Error())
	}
}

func Login(ctx *gin.Context) {
	var userLogin models.UserLogin
	err := ctx.ShouldBind(&userLogin)
	if err != nil {
		logSource.Log.Error("登录信息绑定错误")
		return
	}
	err = ctx.ShouldBind(&userLogin)
	if err != nil {
		logSource.Log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 2,
			"msg":  "用户登陆失败，登陆参数错误",
		})
	}
	id := userLogin.Id
	verifyCode := userLogin.VerifyCode
	password := userLogin.Password
	username := userLogin.UserName
	verifyResult := dataSource.SearchUser(id, username, password, verifyCode)
	if verifyResult.Code == 1 {
		session := sessions.Default(ctx)
		session.Set(username, "success")
		session.Options(sessions.Options{
			MaxAge: int(24 * time.Hour),
		})
		err := session.Save()
		if err != nil {
			logrus.Warn(fmt.Sprintf("%s登录设置session失败，%s", username, err.Error()))
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "登陆成功",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 2,
			"msg":  verifyResult.Msg,
		})
	}
}

package user

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"net/http"
	"otp/dataSource"
	"otp/logSource"
	"otp/models"
	"otp/sendMail"
	"otp/sessionInit"
	"strconv"
	"time"
)

func SendLoginVerifyCode(ctx *gin.Context) {
	username := ctx.PostForm("username")
	user := dataSource.GetUserInfoByUserName(username)
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	randomNumber := rng.Intn(1000000)
	body := fmt.Sprintf("您此次登陆的验证码是：%s\n", strconv.Itoa(randomNumber))
	to := user.Email
	Subject := "登录验证码"
	userId := user.Id
	sendMail.SendEmail(userId, Subject, body, to)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Gexin..950228",
		DB:       0,
	})
	key := user.UserName + "-" + "loginVerifyCode"
	value := strconv.Itoa(randomNumber)
	errSaveVerifyCodeToRedis := rdb.Set(context.Background(), key, value, time.Hour*24).Err()
	if errSaveVerifyCodeToRedis != nil {
		logSource.Log.Error(errSaveVerifyCodeToRedis.Error())
	}
}

func ToLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user/login.html", nil)
}

func Login(ctx *gin.Context) {
	var userLogin models.UserLogin
	userLogin.UserName = ctx.PostForm("username")
	userLogin.Password = ctx.PostForm("password")
	userLogin.SourceUri = ctx.PostForm("redirectUri")
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
	store := sessionInit.InitSession()
	sessions.Sessions("loginSession", store)
	verifyResult := dataSource.SearchUser(id, username, password, verifyCode)
	if verifyResult.Code == 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg": "Success",
			})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "登录验证失败",
			"code": 2,
		})
	}
}

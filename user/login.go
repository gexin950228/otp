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
	"otp/dataSource"
	"otp/logSource"
	"otp/models"
	"otp/sendMail"
	"otp/sessionInit"
	"strconv"
	"time"
)

var session = sessionInit.InitSession()

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
	key := strconv.FormatInt(int64(user.Id), 64) + "-" + "loginVerifyCode"
	value := strconv.FormatInt(int64(randomNumber), 64)
	errSaveVerifyCodeToRedis := rdb.Set(context.Background(), key, value, 120).Err()
	if errSaveVerifyCodeToRedis != nil {
		logSource.Log.Error(errSaveVerifyCodeToRedis.Error())
	}
}

func ToLogin(ctx *gin.Context) {
	redirectUri := ctx.Query("uri")
	ctx.HTML(http.StatusOK, "templates/login.html", gin.H{"Uri": redirectUri})
}

func Login(ctx *gin.Context) {
	uri, _ := ctx.GetPostForm("uri")
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
		session := sessions.Default(ctx)
		session.Set(username, "Success")
		session.Options(sessions.Options{
			MaxAge: int(24 * time.Hour),
		})
		err := session.Save()
		if err != nil {
			logrus.Warn(fmt.Sprintf("%s登录设置session失败，%s", username, err.Error()))
		} else {
			session.Set(username, "Success")
			TokenHandler(username, ctx.Writer, ctx.Request)
			ctx.Redirect(http.StatusOK, uri)
		}

	} else {
		redirectUri := "/response/show/?id=" + username
		ctx.Redirect(http.StatusOK, redirectUri)
	}
}

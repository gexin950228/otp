package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"otp/dataSource"
	"otp/logSource"
	"otp/middleware"
	"otp/models"
	"otp/sendMail"
	"strconv"
	"time"
)

func SendLoginVerifyCode(ctx *gin.Context) {
	username := ctx.PostForm("username")
	user := dataSource.GetUserInfoByUserName(username)
	fmt.Println(user)
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	randomNumber := rng.Intn(1000000)
	body := fmt.Sprintf("您此次登陆的验证码是：%s\n", strconv.Itoa(randomNumber))
	to := user.Email
	Subject := "登录验证码"
	sendMail.SendEmail(Subject, body, to)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Gexin..950228",
		DB:       0,
	})
	key := username
	value := strconv.Itoa(randomNumber)
	fmt.Printf("verifyCode:=================%v\n", value)
	errSaveVerifyCodeToRedis := rdb.Set(context.Background(), key, value, time.Hour*2).Err()
	if errSaveVerifyCodeToRedis != nil {
		logSource.Log.Error(errSaveVerifyCodeToRedis.Error())
	}
	fmt.Println("=============================================")
}

func ToLogin(ctx *gin.Context) {
	uri := ctx.Query("uri")
	fmt.Println(uri)
	ctx.HTML(http.StatusOK, "user/login.html", uri)
}

func Login(ctx *gin.Context) {
	var userLogin models.UserLogin
	err := ctx.ShouldBind(&userLogin)
	if err != nil {
		logSource.Log.Warn(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 2,
			"msg":  "用户登陆失败，登陆参数错误",
		})
	}

	// 验证登录信息
	verifyResult := dataSource.SearchUser(userLogin.UserName, userLogin.Password, userLogin.VerifyCode)
	if verifyResult.Code == 1 {
		claims := middleware.Claims{}
		claims.UserName = userLogin.UserName
		claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
		tokenString := middleware.GenerateToken(claims)
		ctx.SetCookie("otp-token", tokenString, 3600, "/response/search", "www.otp.com", false, true)

		if err != nil {
			logrus.Error(err.Error())
		}

		//ctx.Redirect(http.StatusMovedPermanently, "/response/show/"+userLogin.UserName)
		ctx.Writer.Header().Set("Authorization", tokenString)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  tokenString,
		})
	} else {
		data := map[string]interface{}{
			"code":        "2",
			"redirectUri": userLogin.RedirectUri,
			"msg":         "success",
		}
		ctx.JSON(http.StatusUnauthorized, data)
	}

}

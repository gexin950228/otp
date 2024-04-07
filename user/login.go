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
	"strings"
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
	fmt.Println("============================")
	uri := ctx.Param("uri")
	fmt.Printf("前端拿到的uri: %s\n", uri)
	var toUri string
	if strings.Contains(uri, "search") {
		toUri = "http://127.0.0.1:8080/response/search/"
	} else if strings.Contains(uri, "update") {
		toUri = "http://127.0.0.1:8080/response/update/"
	} else if strings.Contains(uri, "delete") {
		toUri = "http://127.0.0.1:8080/response/delete/"
	} else if strings.Contains(uri, "show") {
		toUri = "http://127.0.0.1:8080/response/show/"
	} else {
		toUri = "http://127.0.0.1:8080/response/add/"
	}
	fmt.Printf("前端传给后端的RedirectUri: %s\n", toUri)
	ctx.HTML(http.StatusOK, "user/login.html", toUri)
}

func Login(ctx *gin.Context) {
	var userLogin models.UserLogin
	fmt.Println(userLogin)
	err := ctx.ShouldBind(&userLogin)
	fmt.Printf("UserLogin: %v\n", userLogin)
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
		fmt.Println("校验成功")
		tokenString := middleware.GenerateToken(userLogin.UserName)
		fmt.Printf("tokenString: %s\n", tokenString)
		ctx.SetCookie("otp-token", tokenString, 3600, "/response/search", "", false, true)

		if err != nil {
			logrus.Error(err.Error())
		}

		//ctx.Redirect(http.StatusMovedPermanently, "/response/show/"+userLogin.UserName)
		ctx.Header("LoginUser", userLogin.UserName)
		ctx.Header(userLogin.UserName+"-Authorization", "Succeed")
		ctx.Header("token", tokenString)
		data := map[string]string{
			"code":           "1",
			"Authentication": "Login Success",
			"msg":            userLogin.RedirectUri,
			"loginUser":      userLogin.UserName,
		}
		ctx.JSON(http.StatusOK, data)
	} else {
		fmt.Println("====================登录校验失败=================")
		data := map[string]interface{}{
			"code":        "2",
			"redirectUri": "http://127.0.0.1:8080/user/login",
			"msg":         "fail",
		}
		ctx.JSON(http.StatusUnauthorized, data)
	}

}

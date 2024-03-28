package dataSource

import (
	"context"
	"fmt"
	"otp/logSource"
	"otp/models"
	"github.com/go-redis/redis/v8"
)

type VerifyResult struct {
	Code int64
	Msg  string
}

func GetUserInfoByUserName(username string) models.UserInfo {
	var user models.UserInfo
	Db.First(&user).Where("username=?", username)
	return user
}

func SearchUser(id int, username string, password string, verifyCode string) VerifyResult {
	var loginInfo models.UserLogin
	Db.First(&loginInfo).Where("username=?", username)
	key := username + "-" + "loginVerifyCode"
	fmt.Println(key)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Gexin..950228",
		DB:       0,
	})
	var verifyResult VerifyResult
	result, err := rdb.Get(context.Background(), key).Result()
	fmt.Println(result)
	if err != nil {
		logSource.Log.Error("登录校验出错，查询校验码出错")
		verifyResult.Code = 2
		verifyResult.Msg = "查询校验码出粗"

	} else {
		if id == loginInfo.Id && username == username && password == password && verifyCode == result {
			verifyResult.Code = 1
			verifyResult.Msg = "校验成功"
		} else {
			verifyResult.Code = 2
			verifyResult.Msg = "提交的信息错误，校验失败。"
		}
	}
	return verifyResult
}

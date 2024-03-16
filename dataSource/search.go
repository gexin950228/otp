package dataSource

import (
	"context"
	"github.com/go-redis/redis/v8"
	"otp/logSource"
	"otp/models"
	"strconv"
)

type VerifyResult struct {
	Code int64
	Msg  string
}

func SearchUser(id int, username string, password string, verifyCode string) VerifyResult {
	var loginInfo models.UserLogin
	Db.First(&loginInfo).Where("id=?", loginInfo.Id)
	key := strconv.FormatInt(int64(loginInfo.Id), 64) + "-" + "loginVerifyCode"
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Gexin..950228",
		DB:       0,
	})
	var verifyResult VerifyResult
	result, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		logSource.Log.Error("登录校验出错，查询校验码出错")
		verifyResult.Code = 2
		verifyResult.Msg = "查询校验码出粗"

	} else {
		if id == loginInfo.Id && username == loginInfo.UserName && password == loginInfo.Password && verifyCode == result {
			verifyResult.Code = 1
			verifyResult.Msg = "校验成功"
		} else {
			verifyResult.Code = 2
			verifyResult.Msg = "提交的信息错误，校验失败。"
		}
	}
	return verifyResult
}

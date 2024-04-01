package dataSource

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"otp/models"
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

func SearchUser(username string, password string, verifyCode string) VerifyResult {
	var loginInfo models.UserInfo
	Db.First(&loginInfo).Where("username=?", loginInfo.UserName)
	fmt.Printf("user in database: %v\n", &loginInfo)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Gexin..950228",
		DB:       0,
	})
	var verifyResult VerifyResult
	fmt.Printf("username to search:%v\n", username)
	result, err := rdb.Get(context.Background(), username).Result()
	fmt.Printf("result inf redis: %v\n", result)
	if err != nil {
		fmt.Println("11111111111111111111111111111111111111")
		fmt.Printf("result: %v\n", result)
		verifyResult.Code = 2
		verifyResult.Msg = "查询校验码出错"

	} else {
		if username == loginInfo.UserName {
			if password == loginInfo.Password {
				if verifyCode == result {
					verifyResult.Code = 1
					verifyResult.Msg = "校验成功"
				} else {
					fmt.Println("验证码错误")
					verifyResult.Code = 2
					verifyResult.Msg = "提交的信息错误，校验失败。"
				}
			} else {
				fmt.Println("3333333333333")
				verifyResult.Code = 2
				verifyResult.Msg = "登录失败"
			}
		} else {
			fmt.Println("用户名错误")
			verifyResult.Code = 2
			verifyResult.Msg = "登录失败"
		}
	}
	return verifyResult
}

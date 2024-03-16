package sessionInit

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"otp/logSource"
	"otp/utils"
)

func InitSession() sessions.Store {
	redisConf := utils.LoadRedisConfig()
	sessionStore, errInitSession := redis.NewStore(redisConf.ConnNum, "tcp", redisConf.Address, redisConf.Password, []byte("secret"))
	if errInitSession != nil {
		logSource.Log.Panic(errInitSession.Error())
		fmt.Println(errInitSession.Error())
		return nil
	} else {
		fmt.Println("初始化session成功")
		logSource.Log.Info("初始化session成功")
		return sessionStore
	}
}

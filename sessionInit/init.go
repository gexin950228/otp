package sessionInit

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"otp/logSource"
)

func InitSession() sessions.Store {
	sessionStore, errInitSession := redis.NewStore(10, "tcp", "localhost:6379", "Gexin..950228", []byte("secret"))
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

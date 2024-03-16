package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type RedisConf struct {
	ConnNum  int    `json:"connNum"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func LoadRedisConfig() *RedisConf {
	var redisConfig RedisConf
	file, errLoadRedis := os.Open("conf/redis.json")
	if errLoadRedis != nil {
		logrus.Error(fmt.Sprintf("加载redis配置文件出错，%s", errLoadRedis.Error()))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	byteData, errReadFile := io.ReadAll(file)
	if errReadFile != nil {
		logrus.Fatal(errReadFile.Error())
	}
	errUnmarshal := json.Unmarshal(byteData, &redisConfig)
	if errUnmarshal != nil {
		logrus.Fatal(errUnmarshal.Error())
	}
	return &redisConfig
}

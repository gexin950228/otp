package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var Log = logrus.New()

func init() {
	logConfig := LoadConfig()
	logDir := filepath.Join(logConfig.LogPath, logConfig.LogName)

	// 设置日志输出文件
	file, err := os.OpenFile(logDir, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("打开日志文件出错")
		return
	}
	Log.Out = file
	levelMapping := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"trace": logrus.TraceLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.FatalLevel,
	}
	Log.SetLevel(levelMapping[logConfig.LogLevel])
}

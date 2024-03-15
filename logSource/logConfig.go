package logSource

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var Log = logrus.New()

func init() {

	logConfig := LoadConfig()
	logPath := logConfig.LogPath
	logName := logConfig.LogName
	log := filepath.Join(logPath, logName)
	file, err := os.OpenFile(log, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	Log.Out = file
	levelMapping := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.DebugLevel,
		"warn":  logrus.WarnLevel,
		"trace": logrus.TraceLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}
	Log.SetLevel(levelMapping[logConfig.LogLevel])
	Log.SetFormatter(&logrus.TextFormatter{})

}

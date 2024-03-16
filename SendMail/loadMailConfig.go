package SendMail

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type MailMsg struct {
	Password string `yaml:"Password"`
	From     string `yaml:"From"`
	Server   string `yaml:"Server"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
}

func LoadMailConfig(filename string) MailMsg {
	mailMsg := MailMsg{}
	file, errOpenFile := os.Open(filename)
	if errOpenFile != nil {
		logrus.Error(fmt.Sprintf("加载配置出错： %s"), errOpenFile.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logrus.Error(fmt.Sprintf("加载配置文件出错， %s", err.Error()))
		}
	}(file)
	byteDate, errReadBytes := io.ReadAll(file)
	if errReadBytes != nil {
		logrus.Error(fmt.Sprintf("加载配置文件出错, %s", errReadBytes.Error()))
	}
	errUnmarshal := json.Unmarshal(byteDate, &mailMsg)
	if errUnmarshal != nil {
		logrus.Error(fmt.Sprintf("写入日志出错, %s", errUnmarshal.Error()))
	}
	return mailMsg
}

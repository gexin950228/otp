package SendMail

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"otp/logSource"
	"otp/utils"
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
		logSource.Log.Error(fmt.Sprintf("加载配置出错： %s"), errOpenFile.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			utils.WriteLog("加载配置文件出错", err.Error())
		}
	}(file)
	byteDate, errReadBytes := io.ReadAll(file)
	if errReadBytes != nil {
		utils.WriteLog("加载配置文件出错", errReadBytes.Error())
	}
	errUnmarshal := json.Unmarshal(byteDate, &mailMsg)
	if errUnmarshal != nil {
		utils.WriteLog("写入日志出错", errUnmarshal.Error())
	}
	return mailMsg
}

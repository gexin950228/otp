package utils

import (
	"os"
	"time"
)

func WriteLog(logSubject string, logContent string) {
	timeNow := time.Now().Format("2006-01-03 15:04:05")
	fileName := "inspection.logSource"
	file, errOpenFile := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0644)
	if errOpenFile != nil {
		WriteLog("日志写入出错", errOpenFile.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	_, err1 := file.WriteString(timeNow + " [" + logSubject + "] " + logContent + "\r\n")
	if err1 != nil {
		WriteLog("日志写入出错", err1.Error())
	}
}

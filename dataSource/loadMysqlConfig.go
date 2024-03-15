package dataSource

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogMode  bool   `json:"logMode"`
}

func LoadMysqlConfig() *MysqlConfig {
	mysqlConfig := MysqlConfig{}
	file, errOpenMysqlConfig := os.Open("conf/mysqlConn.json")
	if errOpenMysqlConfig != nil {
		logrus.Panic(errOpenMysqlConfig.Error())
	}
	defer file.Close()
	byteData, errReadFile := io.ReadAll(file)
	if errReadFile != nil {
		logrus.Fatal(errReadFile.Error())
	}
	errUnmarshal := json.Unmarshal(byteData, &mysqlConfig)
	if errUnmarshal != nil {
		logrus.Fatal(errUnmarshal.Error())
	}
	return &mysqlConfig
}

package logSource

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type LogConfig struct {
	LogPath  string `json:"logPath"`
	LogLevel string `json:"logLevel"`
	LogName  string `json:"logName"`
}

func LoadConfig() *LogConfig {
	logConfig := LogConfig{}
	file, err := os.Open("conf/logConfig.json")
	if err != nil {
		fmt.Printf("Open log config error, error: %v\n", err.Error())
		panic(err)
	}
	defer file.Close()
	byteData, err1 := io.ReadAll(file)
	if err1 != nil {
		return nil
	}
	errUmarshalLogConfig := json.Unmarshal(byteData, &logConfig)
	if errUmarshalLogConfig != nil {
		return nil
	}
	return &logConfig
}

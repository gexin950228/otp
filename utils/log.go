package utils

import (
	"encoding/json"
	"io"
	"os"
)

type LogConfig struct {
	LogPath  string `json:"logPath"`
	LogLevel string `json:"logLevel"`
	LogName  string `json:"logName"`
}

func LoadConfig() *LogConfig {
	LogConfig := LogConfig{}
	file, err := os.Open("conf/logConfig.json")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	byteData, errRead := io.ReadAll(file)
	if errRead != nil {
		return nil
	}
	errUnmarshalLogConfig := json.Unmarshal(byteData, &LogConfig)
	if errUnmarshalLogConfig != nil {
		return nil
	}
	return &LogConfig
}

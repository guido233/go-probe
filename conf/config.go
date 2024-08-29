package conf

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go-probe/logger"
)

var (
	Conf *Config
)

type Websocket struct {
	Host string
	Port string
}

type Config struct {
	ServeName string
	Serial    string
	Websocket Websocket
}

func NewConfig() *Config {
	return &Config{
		ServeName: "",
		Serial:    "",
		Websocket: Websocket{
			Host: "127.0.0.1",
			Port: "3308",
		},
	}
}

func InitConfig() {

	Conf = NewConfig()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Fatalf("Config file not found")
		} else {
			fmt.Println(err.Error())
		}
	}

	err := viper.Unmarshal(Conf)
	if err != nil {
		logger.Fatalf("unable to decode into struct, %v", err)
	}

	// 打印配置信息
	configuration, _ := json.Marshal(Conf)
	logger.Infof("Using conf: %v", string(configuration))
}

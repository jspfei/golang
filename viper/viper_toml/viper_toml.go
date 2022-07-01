package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Redis string
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Port     int
	Host     string
	Username string
	Password string
}

func main() {
	var config Config

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	viper.Unmarshal(&config) //将配置文件绑定到config上
	fmt.Println("config: ", config, "redis: ", config.Redis)
}

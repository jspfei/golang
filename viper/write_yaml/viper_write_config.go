package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.Set("yaml", "this is a example of yaml")
	viper.Set("redis.port", 4405)
	viper.Set("redis.host", "127.0.0.1")

	viper.Set("mysql.port", 3306)
	viper.Set("mysql.host", "192.168.1.0")
	viper.Set("mysql.username", "root")
	viper.Set("mysql.password", "123456")

	if err := viper.WriteConfig(); err != nil {
		fmt.Println(err)
	}
}

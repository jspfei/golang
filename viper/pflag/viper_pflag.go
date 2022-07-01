package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.Int("redis.port", 3302, "redis port")

	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println()
		return
	}
	host := viper.Get("mysql.host")
	username := viper.GetString("mysql.username")
	port := viper.GetInt("mysql.port")
	redisHost := viper.GetString("redis.host")
	redisPort := viper.GetInt("redis.port")

	fmt.Println("mysql - host: ", host, ", username: ", username, ", port:", port)
	fmt.Println("redis - host: ", redisHost, ", port : ", redisPort)
}

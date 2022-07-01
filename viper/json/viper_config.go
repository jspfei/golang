package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() //根据上面配置加载文件
	if err != nil {
		fmt.Println(err)
		return
	}
	host := viper.Get("mysql.host")
	username := viper.GetString("mysql.username")
	port := viper.GetInt("mysql.port")
	portsSlice := viper.GetIntSlice("mysql.ports")
	metricPort := viper.GetInt("mysql.port")
	redis := viper.Get("redis")

	mysqlMap := viper.GetStringMapString("mysql")

	if viper.IsSet("mysql.host") {
		fmt.Println("[IsSet()]mysql.host is set")
	} else {
		fmt.Println("[IsSet()]mysql.host is not set")
	}

	fmt.Println("mysql - host: ", host, ", username: ", username, ", port:", port)

	fmt.Println("mysql ports :", portsSlice)
	fmt.Println("metric port: ", metricPort)
	fmt.Println("redis - ", redis)
	fmt.Println("mysqlmap - ", mysqlMap, ", username: ", mysqlMap["username"])
}

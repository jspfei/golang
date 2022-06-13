package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Run(configName string) error {
	config := Config{
		Name: configName,
	}
	if err := config.init(); err != nil {
		return err
	}
	config.watcher()

	return nil
}

func (config *Config) init() error {
	if config.Name != "" {
		viper.SetConfigFile(config.Name)
	} else {
		//配置人家路径 ./config.yaml
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file : %s\n", err))
	}
	return nil
}

func (config *Config) watcher() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("Config file changed,[%s]", in.Name)
	})
}

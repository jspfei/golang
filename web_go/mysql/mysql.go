package mysql

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var once sync.Once
var db *gorm.DB

// Get 获取MySQL连接实例
func Get() *gorm.DB {
	once.Do(func() {
		db = initialize()
	})
	return db
}

func initialize() *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=%s",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.name"),
		viper.GetString("mysql.charset"))
	db, err := gorm.Open("mysql", config)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(viper.GetInt("mysql.max-idle-connections"))
	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max-open-connections"))
	db.LogMode(viper.GetBool("mysql.log-mode"))
	return db
}

// Close 关闭MySQL连接
func Close(db *gorm.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

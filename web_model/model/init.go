package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-dirver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB

func Init() error {
	var err error

	DB, err = sql.Open("mysql", viper.GetString("mysql.source_name"))
	if nil != err {
		return err
	}

	DB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))

	err = DB.Ping()
	if err != nil {
		return err
	} else {
		log.Println("Mysql Startup Normal!")
	}
	return nil
}

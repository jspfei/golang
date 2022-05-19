package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	User_id  int    `db : "user_id"`
	Username string `db : "username"`
	Password string `db : "password"`
}

var DB *sqlx.DB

func check(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/school")
	check(err)
	DB = database
}

func main() {
	_, err := DB.Exec("update user_login set password=? where user_id=?", "246810", 10002)
	check(err)

	fmt.Println("update success")
}

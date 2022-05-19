package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	// TAG 原信息就是表中定义的属性名，db 是随意的
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
	// 连接数据库 --- "用户名:密码@协议(地址:端口)/数据库名"
	// 生产环境密码单独写在配置文件中，使用 MD5 加密
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/school")
	check(err)

	DB = database
}

func main() {
	dsn := "insert into user_login (user_id,username,password) values(?,?,?)"
	result, err := DB.Exec(dsn, 10003, "boy", "by123456")
	check(err)

	id, err := result.LastInsertId()

	check(err)

	fmt.Println("data insert success", id)
}

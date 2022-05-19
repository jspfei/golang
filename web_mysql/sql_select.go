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

// 数据库结构体
var DB *sqlx.DB

func init() {
	// 连接数据库 --- "用户名:密码@协议(地址:端口)/数据库名"
	// 生产环境密码单独写在配置文件中，使用 MD5 加密
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/school")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	DB = database

}

func main() {
	// 定义获取数据的切片，用来装载查询结果
	var users []User
	err := DB.Select(&users, "select * from user_login")
	if err != nil {
		fmt.Println("查询数据失败,", err)
		return
	}
	fmt.Println("查询的数据:", users)
}

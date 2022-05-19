package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// 数据库结构体
var DB *sqlx.DB

//创建需要返回值的结构体
type User struct {
	User_id  int    `db : "user_id"`
	Username string `db : "username"`
	Password string `db : "password"`
}

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

//返回一个多层的map
func getsimple() []User {

	// //关闭数据库连接
	// defer DB.Close()
	//对sql执行后的每一行进行处理
	var users []User
	err1 := DB.Select(&users, "select * from user_login")
	check(err1)

	fmt.Println("执行了", users)
	//返回多层map
	return users
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			//装载需要返回的json
			"data":    getsimple(),
			"status":  200,
			"message": "您成功获取到数据",
		})
	})

	r.Run(":8888")
}

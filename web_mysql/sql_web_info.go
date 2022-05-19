package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const form = `<html>
	<head></head>
	<body>
	<form action="#" method="post" name="bar">
		<input type="test" name="in"/>
		<input type="submit" value="search"/>
	</form>
	</body>
</html>`

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

func formServer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	switch req.Method {
	case "GET":
		io.WriteString(res, form)
	case "POST":
		req.ParseForm()
		id, err := strconv.Atoi(req.Form["in"][0])
		check(err)

		var user []User
		err = DB.Select(&user, "select * from user_login where user_id=?", id)
		check(err)

		for _, v := range user {
			io.WriteString(res, strconv.Itoa(v.User_id))
			io.WriteString(res, "\n")
			io.WriteString(res, v.Username)
			io.WriteString(res, "\n")
			io.WriteString(res, v.Password)
			io.WriteString(res, "\n")
		}
	}
}

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/school")
	check(err)
	DB = database
}

func main() {
	http.HandleFunc("/mysql", formServer)
	if err := http.ListenAndServe("127.0.0.1:8888", nil); err != nil {
		fmt.Println("启动监听失败，err:", err)
		return
	}
}

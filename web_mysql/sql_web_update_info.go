package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const form = `<!doctype html>
<html lang="en">
	<head>
	<meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
 
	</head>
	<body>
	<form action="#" method="post" name="bar">
	    <p>user_id	<input type="test" name="in"/>  </p>
		<p>username <input type="test" name="in"/> </p>
		<p>password <input type="test" name="in"/> </p>
		<input type="submit" value="update"/>
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
		id, _ := strconv.Atoi(req.Form["in"][0])

		ps := req.Form["in"][2]

		_, err := DB.Exec("update user_login set password=?  where user_id=?", ps, id)
		if err != nil {
			fmt.Println("update data err:", err)
			return
		} else {
			io.WriteString(res, "update succ")
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

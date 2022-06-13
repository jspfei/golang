package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Transaction(db *sql.DB) {
	//开启事务
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	result, err := tx.Exec("insert into user(id,user_name,password)values(?,?,?)", 2, "js", "12345")
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	fmt.Println("result", result)

	exec, err := tx.Exec("update user set password=? where id=?", "admin", 1)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	fmt.Println("exec", exec)

	//提交事务
	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		panic(err)
	}
}

func main() {
	dns := "root:123456@tcp(127.0.0.1:3306)/school"

	db, err := sql.Open("mysql", dns)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	Transaction(db)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Title string
}

var myTemplate *template.Template

func userInfo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("user login")
	fmt.Fprintf(res, "<h1>成功登陆</h1>")

	p := Person{
		Name:  "ssss",
		Age:   40,
		Title: "个人网",
	}

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("文件加载失败，err:", err)
		return
	}

	myTemplate.Execute(file, p)
}

func initTemplate(filename string) (err error) {
	myTemplate1, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	myTemplate = myTemplate1
	return
}

func main() {
	initTemplate("index.html")

	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

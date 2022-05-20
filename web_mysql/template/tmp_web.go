package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int

	Title string
}

func main() {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	p := Person{Name: "jj", Age: 23, Title: "经济技术"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("This is error ", err.Error())
	}
}

package main

import (
	"html/template"
	"log"
	"net/http"
)

type Address struct {
	Province string
	City     string
}

type User struct {
	Name string
	Age  int
	Addr Address
}

func index(w http.ResponseWriter, r *http.Request) {
	//解析指定文件生成的模板对象
	tmpl, err := template.ParseFiles("./public/index.html")
	if err != nil {
		log.Fatal(err)
	}
	user := User{
		Name: "彭于晏",
		Age:  28,
		Addr: Address{Province: "台湾省", City: "澎湖县"},
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	files := http.FileServer(http.Dir("./public"))
	http.Handle("/", files)
	http.HandleFunc("/index", index)

	err := http.ListenAndServe("127.0.0.1:2018", nil)
	log.Fatal(err)
}

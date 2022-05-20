package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name  string
	Age   int
	Title string
}

// 模板变量
var myTemplate *template.Template

func userInfo(res http.ResponseWriter, req *http.Request) {
	fmt.Print("user login")
	fmt.Fprintf(res, "<h1>成功登录</h1>")
	var arr []Person
	p1 := Person{Name: "zhangsan", Age: 18, Title: "个人网站"}
	p2 := Person{Name: "lisi", Age: 20, Title: "个人网站"}
	p3 := Person{Name: "wangwu", Age: 22, Title: "个人网站"}
	arr = append(arr, p1)
	arr = append(arr, p2)
	arr = append(arr, p3)

	myTemplate.Execute(res, arr)

}

// 模板函数
func initTemplate(filename string) (err error) {
	// 读取模板文件
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("模板加载失败：", err)
		return
	}
	return
}

func main() {
	// 初始化模板
	initTemplate("index1.html")
	// 路由到站点，调用业务处理函数
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("HTTP LISTEN FAILED,err:", err)
		return
	}
}

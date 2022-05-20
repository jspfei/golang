package main

import (
	"fmt"
	"io"
	"net/http"
)

const form = `<html><body>

</body></html>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello world!</h1>")
	panic("test panic")
}

func FormServer(res http.ResponseWriter, request *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	// 提交类型判断
	switch request.Method {
	case "GET":
		io.WriteString(res, form)
	case "POST":
		request.ParseForm()
		// 读取第一个框体内容给 res 做页面回应
		// 若获取第二个文本框内容则使用 request.Form["in"][1]
		io.WriteString(res, request.Form["in"][0])
		io.WriteString(res, "\n")
		io.WriteString(res, request.FormValue("in"))
	}

}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				fmt.Printf("[%v]caught panic:%v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe("127.0.0.1:8888", nil); err != nil {
		fmt.Println("服务启动失败，err", err)
		return
	}
}

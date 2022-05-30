package main

import (
	"fmt"
	"net"
)

func main() {
	//拨号建立与服务器的连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")

	if err != nil {
		fmt.Println("连接服务器失败，err:", err)
		return
	}

	defer conn.Close()
	//向 server 发送信息
	_, err = conn.Write([]byte("hello, tom"))
	if err != nil {
		fmt.Println("发送信息失败，err:", err)
		return
	}
}

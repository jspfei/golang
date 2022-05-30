package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	var buf [1024]byte
	n, err := conn.Read(buf[:])

	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	str := string(buf[:n])
	fmt.Println("收到客户端发来的数据：", str)
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:20000")

	if err != nil {
		fmt.Println("启动服务器失败，err:", err)
		return
	}

	defer listener.Close()

	for {
		//建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受客户端连接失败，err:", err)
			continue
		}

		//启动一个 goroutine 处理客户端连接
		go process(conn)
	}
}

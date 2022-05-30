package main

import (
	"fmt"
	"net"
)

func process(listener net.UDPConn) {
	defer listener.Close()

	for {
		var buf [1024]byte
		n, addr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("接收信息失败，err:", err)
			return
		}
		fmt.Printf("接收到来自 %v 的消息:%v\n", addr, string(buf[:n]))

		n, err = listener.WriteToUDP([]byte("hi"), addr)
		if err != nil {
			fmt.Println("回复失败,err:", err)
			return
		}
	}
}

//udp服务
func main() {
	//1.监听
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 30000,
	})

	if err != nil {
		fmt.Println("启动 server 失败，err:", err)
		return
	}

	process(*listener)
}

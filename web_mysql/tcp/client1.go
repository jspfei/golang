package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	for {

		conn, err := net.Dial("tcp", "127.0.0.1:20000")

		if err != nil {
			fmt.Println("连接服务器失败：err:", err)
			return
		}

		defer conn.Close()

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取信息失败，err:", err)
			return
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("发送信息失败，err:", err)
			return
		}
	}

}

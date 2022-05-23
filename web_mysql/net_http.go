package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.taobao.com",
	"http://www.baidu.com",
	"http://www.google.com",
}

func main() {
	for _, v := range url {
		client := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}

		req, err := client.Head(v)
		if err != nil {
			fmt.Println("获取请求失败，err:", err)
		}

		fmt.Println("来自 %s 的网站，状态是 %s\n", v, req.Status)
	}
}

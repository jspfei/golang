package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello world \n this is test"

	err := ioutil.WriteFile("./tmp1.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("文件写入错误", err)
		return
	}
}

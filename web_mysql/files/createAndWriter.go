package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("tmp.txt", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}

	defer file.Close()

	str := "hello world"
	file.Write([]byte("this is test\n"))
	file.WriteString(str)
}

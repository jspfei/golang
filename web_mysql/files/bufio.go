package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("tmp2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}

	defer file.Close()
	write := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		write.WriteString("hello world \n")
	}
	write.Flush()
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./abc.txt")

	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	count := 0

	for {
		str, _, err := reader.ReadLine()
		count++

		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("读取文件内容失败", err)
			return
		}

		if count%2 == 1 {
			fmt.Println(string(str))
		}
	}
}

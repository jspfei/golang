package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile() {
	file, err := os.Open("./abc.txt")

	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}

	defer file.Close()
	fmt.Println("文件打开成功")

	//读取
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')

		if err == io.EOF {
			fmt.Println(str)
			return
		}

		if err != nil {
			fmt.Println("read from file failed,err", err)
			return
		}

		fmt.Print(str)
	}

}

func main() {
	readFile()
}

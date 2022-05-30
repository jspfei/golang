package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Read file failed,err:", err)
		return
	}

	fmt.Println(string(content))
}

func main() {
	readFile("./abc.txt")
}

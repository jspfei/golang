package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello 函数")
}

func main() {
	defer fmt.Println("defer3")
	go hello()

	time.Sleep(time.Second)
	fmt.Println("main 函数")
	defer fmt.Println("defer2")
	defer fmt.Println("defer1")
}

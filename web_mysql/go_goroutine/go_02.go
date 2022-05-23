package main

import (
	"fmt"
	"sync"
)

var sw sync.WaitGroup

func hello() {
	fmt.Println("hello 函数")
	sw.Done()
}

func test() {
	fmt.Println("test 函数")
	sw.Done()
}

func main() {
	defer fmt.Println("defer 语句")

	sw.Add(2)
	go hello()
	go test()

	fmt.Println("main 函数")
	sw.Wait()

}

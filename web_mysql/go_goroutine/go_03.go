package main

import (
	"fmt"
	"sync"
)

var sw sync.WaitGroup

func hello(i int) {
	fmt.Println("hello 函数", i)
	if i == 8 {
		panic("宕机报警")
	}
	sw.Done()
}

func main() {
	defer fmt.Println("defer 语句")
	sw.Add(10)
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	fmt.Println("main 函数")
	sw.Wait()
}

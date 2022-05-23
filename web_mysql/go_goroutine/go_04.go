package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var sw sync.WaitGroup

func a() {
	defer sw.Done()

	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	sw.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	// 使用 1 个逻辑核心数跑 Go 程序
	runtime.GOMAXPROCS(1)
	sw.Add(2)
	go a()
	go b()
	time.Sleep(time.Second)
	sw.Wait()
}

package main

import (
	"fmt"
	"sync"
)

var x int64

var sw sync.WaitGroup

var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	sw.Done()
}

func main() {
	sw.Add(2)
	go add()
	go add()
	sw.Wait()
	fmt.Println(x)
}

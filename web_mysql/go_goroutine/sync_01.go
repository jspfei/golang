package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var m = sync.Map{}

func demo() {
	sw := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		sw.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("K:%s; V : %v \n", key, value)
			sw.Done()
		}(i)
	}
	sw.Wait()
}

func tickDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

func main() {
	tickDemo()
}

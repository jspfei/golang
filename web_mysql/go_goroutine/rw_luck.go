package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var sw sync.WaitGroup

var lock sync.Mutex

var rwlock sync.RWMutex

func read() {
	defer sw.Done()
	rwlock.RLock()

	time.Sleep(time.Millisecond * 1)
	rwlock.RUnlock()
}

func write() {
	defer sw.Done()

	rwlock.Lock()

	x++

	time.Sleep(time.Millisecond * 50)
	rwlock.Unlock()
}

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		sw.Add(1)
		go read()
	}

	sw.Wait()
	end := time.Now()
	fmt.Printf("用时: %v .\n", end.Sub(start))

}

package singleton

import (
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(parCount)

	instance := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instance[index] = GetInstance()
			wg.Done()

		}(i)
	}

	//关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instance[i] != instance[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}

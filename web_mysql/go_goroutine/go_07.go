package main

import (
	"fmt"
	"math"
	"time"
)

var ch1 = make(chan string, 100)
var ch2 = make(chan string, 100)

func sendK1(ch chan string) {
	for i := 0; i < math.MaxInt64; i++ {
		ch <- fmt.Sprintf("k1:%d", i)
		time.Sleep(time.Millisecond * 50)
	}
}
func sendK2(ch chan string) {
	for i := 0; i < math.MaxInt64; i++ {
		ch <- fmt.Sprintf("k2:%d", i)
		time.Sleep(time.Millisecond * 50)
	}
}
func main() {
	go sendK1(ch1)
	go sendK2(ch2)

	for {
		select {
		case ret := <-ch1:
			fmt.Println(ret)
		case ret := <-ch2:
			fmt.Println(ret)
		default:
			fmt.Println("没有取到数据")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

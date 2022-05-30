package main

import "fmt"

var ch chan int

func writeCh(ch chan<- int) {
	ch <- 10
}

func readCh(ch <-chan int) int {
	ret := <-ch
	return ret
}

func main() {
	ch = make(chan int, 1)
	writeCh(ch)
	fmt.Println(readCh(ch))
}

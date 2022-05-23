package main

import "fmt"

func main() {
	var ch1 chan int
	var ch2 chan bool

	fmt.Println("ch1:", ch1)
	fmt.Println("ch2:", ch2)

	ch3 := make(chan int, 2)

	ch3 <- 10
	ch3 <- 20

	result := <-ch3
	result2 := <-ch3
	fmt.Println(result, result2)

	// result3 := <-ch3
	// fmt.Println(result3)
}

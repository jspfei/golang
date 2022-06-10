package main

import "fmt"

type cback func(int) int

func main() {

	test_cback(1, callback)
	test_cback(2, func(x int) int {
		fmt.Printf("回调：x : %d \n", x)
		return x
	})

	x, y := 1, 2
	fmt.Println(test(x, y, add))
}

func test_cback(x int, f cback) {
	f(x)

}

func callback(x int) int {
	fmt.Printf("回调：x : %d \n", x)
	return x
}

type Callback func(x, y int) int

func add(x, y int) int {
	return x + y
}

func test(x, y int, callback Callback) int {
	return callback(x, y)
}

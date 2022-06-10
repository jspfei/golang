package main

import "fmt"

func Oper(n int) int {
	if n < 2 {
		return n
	}

	return Oper(n-2) + Oper(n-1)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", Oper(i))
	}
}

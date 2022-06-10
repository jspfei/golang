package main

import "fmt"

func Oper(n uint64) (result uint64) {
	if n > 0 {
		result = n * (n - 1)
		return result
	}

	return result
}

func main() {
	var i int = 5
	fmt.Printf("%d 的阶乘是 %d \n", i, Oper(uint64(i)))
}

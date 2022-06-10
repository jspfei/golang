package main

import (
	"fmt"

	"math"
)

func main() {
	var t1 int = 100
	var t2 int = 200
	var result int
	result = max(t1, t2)
	fmt.Println("max: ", result)

	a, b := multi_value(3, 5)
	fmt.Println("和为：", a, "积为：", b)

	getRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getRoot(9))
}

func multi_value(num1, num2 int) (int, int) {
	result1 := num1 + num2
	result2 := num1 * num2
	return result1, result2
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

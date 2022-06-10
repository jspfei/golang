package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println("元素和为：", sum(numbers))
}

func sum(arr [5]int) int {
	s := 0

	for i := range arr {
		s += arr[i]
	}

	return s

}

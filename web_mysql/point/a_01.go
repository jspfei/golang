package main

import (
	"fmt"
)

func main() {

	var balance = [5]float32{1000.0, 2.0, 3.0, 7.0, 50.0}

	var test1 = [...]int{1, 2, 3, 4}

	fmt.Println(test1)
	fmt.Println(len(test1))

	test1 = [4]int{1, 2, 3, 4}

	fmt.Println(test1)
	var a int = test1[1]
	fmt.Println(a)

	var Array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		fmt.Println(Array[i])
	}

	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i + 100
	}

	for j := 0; j < 10; j++ {
		fmt.PPrintf("Element[%d] = %d\n", j, a[j])
	}

	var arr = []int{1, 3, 5, 7, 9}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

}

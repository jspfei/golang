package main

import "fmt"

func a1() {
	var test = [5]int{1, 2, 3, 4, 5}
	fmt.Println(test)
	var a int = test[1]
	fmt.Println(a)
}

func a2() {
	var Array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 10; i++ {
		fmt.Println(Array[i])
	}
}

func a3() {
	var Array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for i := 0; i < len(Array); i++ {
		fmt.Print(Array[i])
	}
	fmt.Println("数组长度：", len(Array))
}

func main() {
	a1()
	a2()

	a3()
}

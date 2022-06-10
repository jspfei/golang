package main

import "fmt"

func main() {
	strings := []string{"google", "runoob"}

	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 4, 5}
	for i, x := range numbers {
		fmt.printf("第 %d 位 x 的值 = %d \n", i, x)
	}
}

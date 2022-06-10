package main

import "fmt"

func main() {
	var grade int

	fmt.Println("请输入分数：")
	fmt.Scanf("%d", &grade)

	a := grade

	if a >= 90 {
		fmt.Println("成绩为A")
	} else {
		fmt.Println("成绩为B")
	}
}

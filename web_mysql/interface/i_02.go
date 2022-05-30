package main

import "fmt"

type Test interface{}

func main() {
	var t Test
	var a interface{}
	var b int
	a = b
	fmt.Printf("a 数据类型：%T\n", a)
	fmt.Printf("Test 数据类型:%T\n", t)
}

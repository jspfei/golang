package main

import (
	"fmt"
	"reflect"
)

func test(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println(t)

	v := reflect.ValueOf(i)
	fmt.Println(v)
}

func main() {
	var a int = 10
	test(a)
}

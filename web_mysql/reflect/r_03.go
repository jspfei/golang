package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func prints(i int) string {
	fmt.Println("i=", i)
	return strconv.Itoa(i)
}

func main() {
	fv := reflect.ValueOf(prints)

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(20)

	result := fv.Call(params)
	fmt.Println("reuslt 类型：%T\n", result)
	fmt.Printf("result 转换后类型:%T,值是:%s\n", result[0].Interface().(string), result[0].Interface().(string))
}

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func test(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("类型：", t)
	v := reflect.ValueOf(i)
	k := v.Kind()
	fmt.Println("类别：", k)

	// 转换成接口类型

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Println("结构： %v 类型：%T\n", stu, stu)
	}

	num := v.NumField()
	fmt.Println("字段数量：", num)

	numOfMethod := v.NumMethod()
	fmt.Println("方法数量;", numOfMethod)
}

func main() {
	var stu Student = Student{
		Name:  "ss",
		Age:   12,
		Score: 99,
	}
	test(stu)
}

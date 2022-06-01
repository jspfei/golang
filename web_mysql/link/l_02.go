package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student // 单链，下一个
}

func main() {
	// 头部结构节点
	var head Student
	head.Name = "head"
	head.Age = 18
	head.Score = 100

	// 第二个结构体节点
	var stu1 Student
	stu1.Name = "stu1"
	stu1.Age = 20
	stu1.Score = 90

	// 第一个指针
	head.next = &stu1

	// 第三个结构体节点
	var stu2 Student
	stu2.Name = "stu2"
	stu2.Age = 30
	stu2.Score = 80

	// 第二个指针
	stu1.next = &stu2

	// 调用遍历链表函数，传第一个结构体地址
	Req(&head)
}

// 遍历链表
func Req(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

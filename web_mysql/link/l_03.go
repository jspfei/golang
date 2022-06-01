package main

import (
	"fmt"
	"math/rand"
)

type Students struct {
	Name  string
	Age   int
	Score float32
	next  *Students
}

func main() {
	var stu1 Students
	stu1.Name = "stu1"
	stu1.Age = 25
	stu1.Score = 62

	var stu2 Students
	stu2.Name = "stu2"
	stu2.Age = 30
	stu2.Score = 100

	stu1.next = &stu2

	// 声明尾部变量
	var tail = &stu2
	// 尾部批量添加节点
	for i := 3; i < 10; i++ {
		// 节点定义
		var stu Students = Students{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		// 生成结构体串联
		tail.next = &stu
		tail = &stu
	}

	Req(&stu1)
}

func Req(p *Students) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

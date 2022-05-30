package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
}

type Test interface {
	Print()
	Sleep()
}

func (p *Student) Print() {
	fmt.Println("name:", p.Name)
	fmt.Println("age:", p.Age)
	fmt.Println("score:", p.Score)
}

func (p *Student) Sleep() {
	fmt.Println("学生在睡觉")
}

func main() {
	var t Test

	var stu Student = Student{
		Name:  "jf",
		Age:   18,
		Score: 90,
	}

	t = &stu

	t.Print()
	t.Sleep()
}

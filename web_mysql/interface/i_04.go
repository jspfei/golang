package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
}

type Teacher struct {
	Name  string
	Age   int
	Class int
}

type Test interface {
	Print()
	Sleep()
}

func (p Student) Print() {
	fmt.Println("name:", p.Name)
	fmt.Println("age:", p.Age)
	fmt.Println("score:", p.Score)
}

func (p Student) Sleep() {
	fmt.Println("学生在睡觉")
}

func (t Teacher) Print() {
	fmt.Println("name:", t.Name)
	fmt.Println("age:", t.Age)
	fmt.Println("class:", t.Class)
}

func (t Teacher) Sleep() {
	fmt.Println("教师在休息")
}

func main() {
	var t Test

	var stu Student = Student{
		Name:  "jf",
		Age:   18,
		Score: 90,
	}

	var tea Teacher = Teacher{
		Name:  "lisai",
		Age:   50,
		Class: 01,
	}

	t = stu
	t.Print()
	t.Sleep()
	fmt.Println("--------------------")
	t = tea
	t.Print()
	t.Sleep()
}

package main

import "fmt"

type Persion struct {
	Name string
	Age  int
}

type Student struct {
	Persion
	Score float32
}

type Teacher struct {
	Persion
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

	var stu Student
	stu.Name = "jf"
	stu.Age = 18
	stu.Score = 90

	var tea Teacher
	tea.Name = "lisi"
	tea.Age = 60
	tea.Class = 01

	t = stu
	t.Print()
	t.Sleep()
	fmt.Println("-------------")
	t = tea
	t.Print()
	t.Sleep()

}

package main

import "fmt"

type Test1 interface {
	Print()
}

type Test2 interface {
	Sleep()
}

type Student struct {
	Name string
	Age  int
}

func (s Student) Print() {
	fmt.Println("name:", s.Name)
	fmt.Println("age:", s.Age)
}

func (s Student) Sleep() {
	fmt.Println("正在睡觉")
}

func main() {
	var t1 Test1
	var t2 Test2

	var stu Student = Student{
		Name: "sssd",
		Age:  18,
	}

	t1 = stu
	t2 = stu
	t1.Print()
	t2.Sleep()
}

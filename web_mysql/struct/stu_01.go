package main

import "fmt"

type integer int

func d1() {
	var i integer = 1000

	fmt.Println(i)

	var j int = 100

	j = int(i)
	fmt.Println(j)
}

type Student struct {
	Number int
}

type Stu Student

func d2() {
	var a Student
	a = Student{30}

	var b Stu
	b = Stu{300}

	a = Student(b)

	fmt.Println(a)
}

func main() {
	d2()
}

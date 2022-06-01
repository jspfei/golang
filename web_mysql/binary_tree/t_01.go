package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32

	left  *Student
	right *Student
}

func main() {
	var root Student
	root.Name = "root"
	root.Age = 10
	root.Score = 80

	var left1 Student
	left1.Name = "left1"
	left1.Age = 20
	left1.Score = 70

	root.left = &left1

	var right1 Student
	right1.Name = "right2"
	right1.Age = 60
	right1.Score = 60

	root.right = &right1

	var left2 Student
	left2.Name = "left2"
	left2.Age = 40
	left2.Score = 50

	left1.left = &left2

	ForwardReq(&root)
	MidReq(&root)
	BackReq(&root)
}

func ForwardReq(p *Student) {
	if p == nil {
		return
	}

	fmt.Println(p)
	Req(p.left)
	Req(p.right)
}

func MidReq(p *Student) {
	if p == nil {
		return
	}
	Req(p.left)
	fmt.Println(p)

	Req(p.right)
}

func BackReq(p *Student) {
	if p == nil {
		return
	}

	fmt.Println(p)
	Req(p.left)
	Req(p.right)
}

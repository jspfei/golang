package main

import "fmt"

type Integer int

func (i Integer) convert() string {
	return fmt.Sprintf("%d", i)
}

func (p *Integer) set(b Integer) {
	*p = b
}

func main() {
	var i Integer
	i = 100

	s := i.convert()
	fmt.Printf("类型：%T，值：%s\n", s, s)
	fmt.Printf("类型：%T，值：%d\n", i, i)

	i.set(200)
	fmt.Println()
}

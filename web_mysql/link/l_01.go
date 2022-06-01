package main

import "fmt"

type Car struct {
	Name  string
	Color string
	next  Person
}

type Person struct {
	Owner string
	Age   uint
}

func main() {
	var c Car
	c.Name = "baoma"
	c.Color = "yellow"
	c.next.Owner = "zc"
	c.next.Age = 25
	fmt.Println(c)
}

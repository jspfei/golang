package main

import "fmt"

type Car struct {
	Name  string
	Color string
}

func main() {
	var car1 = new(Car)
	fmt.Println(car1)

	var car2 = NewCar("ts", "red")
	fmt.Println(car2)
}

func NewCar(name string, color string) *Car {
	return &Car{
		Name:  name,
		Color: color,
	}
}

package main

import "fmt"

type Car struct {
	weight int
	name   string
}

func (c *Car) Run() {
	fmt.Println("Running")
}

type Bike struct {
	c     Car
	wheel int
}

type Train struct {
	Car
	wheel int
}

func main() {
	var bike Bike
	bike.c.weight = 100
	bike.c.name = "bike"
	bike.wheel = 2

	var train Train
	train.weight = 140000
	train.name = "train"
	train.wheel = 8

	fmt.Println(bike)

	bike.c.Run()

	fmt.Println(train)

	train.Run()

}

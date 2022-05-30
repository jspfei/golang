package main

import "fmt"

type Car struct {
	weight int
	name   string
}

func (c Car) Run() {
	fmt.Println("Running")
}

type Train struct {
	Car
	wheel int
}

func (p Train) String() string {
	str := fmt.Sprintf("name=[%s],wheel=[%d]", p.name, p.wheel)
	return str
}

func main() {
	var train Train
	train.weight = 140000
	train.name = "train"
	train.wheel = 9

	fmt.Println(train)
	train.Run()
	fmt.Printf("%s", train)
}

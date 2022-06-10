package main

import "fmt"

type Car struct {
	Name  string
	Color string
}

func (c Car) Call() {
	fmt.Printf("%s 品牌的汽车，颜色是%s正在鸣笛", c.Name, c.Color)
}

func (c Car) Run() {
	fmt.Printf("%s 品牌的汽车，颜色是%s正在行驶", c.Name, c.Color)
}

func main() {
	c1 := new(Car)
	c1.Name = "奔驰"
	c1.Color = "黑色"
	c1.Call()
	fmt.Println()

	c2 := new(Car)
	c2.Name = "宝马"
	c2.Color = "白色"
	c2.Run()
	fmt.Println()
}

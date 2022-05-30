package main

import "fmt"

type Computer struct {
	Brand string  //品牌
	Price float32 //价格
}

type USB interface {
	mouse()
	store()
	fan()
}

func (c Computer) mouse() {
	fmt.Println("鼠标控制")
}

func (c Computer) store() {
	fmt.Println("U 盘存储")
}

func (c Computer) fan() {
	fmt.Println("吹风扇")
}

func main() {
	var u USB

	var computer Computer
	computer.Brand = "thinkpad"
	computer.Price = 5000

	u = computer
	u.mouse()
	u.store()
	u.fan()
}

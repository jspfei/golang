package main

import "fmt"

type FuncType func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func Calc(a, b int, test FuncType) (result int) {
	fmt.Println("Calc")
	result = test(a, b)
	return
}

func main() {
	x := Calc(4, 2, Add)
	y := Calc(4, 2, Minus)
	z := Calc(4, 2, Mul)
	k := Calc(4, 2, Div)

	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("z =", z)
	fmt.Println("k =", k)

}

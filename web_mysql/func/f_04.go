package main

import "fmt"

func closePackage() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := closePackage()
	fmt.Println(nextNumber()) // 1
	fmt.Println(nextNumber()) // 2
	fmt.Println(nextNumber()) // 3

}

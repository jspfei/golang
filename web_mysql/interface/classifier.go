package main

import (
	"fmt"
)

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("%d is a bool\n", i)
		case int, int32, int64:
			fmt.Printf("%d is int\n", i)
		case float32, float64:
			fmt.Printf("%d is a float\n", i)
		case string:
			fmt.Printf("%d is a string\n", i)
		default:
			fmt.Printf("unkown")
		}
	}
}

func main() {
	classifier("jf", 18, 90.5, false, nil)
}

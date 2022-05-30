package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

type Student1 struct {
	Name  string  `json:"stu_name"`
	Age   int     `json:"stu_age"`
	Score float32 `json:"stu_score"`
}

func main() {
	var stu Student1 = Student1{
		Name:  "stu1",
		Age:   18,
		Score: 80,
	}

	data, err := json.Marshal(stu)

	if err != nil {
		fmt.Println("json error:", err)
		return
	}

	fmt.Println(data)
	fmt.Println(string(data))
}

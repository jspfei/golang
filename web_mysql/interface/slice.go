package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

type StudentArray []Student

func (sa StudentArray) Len() int {
	return len(sa)
}

func (sa StudentArray) Less(i, j int) bool {
	return sa[i].Name > sa[j].Name
}

func (sa StudentArray) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}

func main() {
	var stus StudentArray

	for i := 0; i < 10; i++ {
		var stu Student = Student{
			Name:  fmt.Sprintf("stu%d", rand.Intn(100)),
			Age:   rand.Intn(120),
			Score: rand.Float32() * 100,
		}

		stus = append(stus, stu)
	}

	for _, v := range stus {
		fmt.Println(v)
	}

	fmt.Println("-------------------------")

	sort.Sort(stus)

	for _, v := range stus {
		fmt.Println(v)
	}
}

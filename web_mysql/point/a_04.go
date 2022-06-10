package main

import "fmt"

func m1() {
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i + 100
	}

	for j := 0; j < 10; j++ {
		fmt.Printf("Element [%d] = %d\n", j, a[j])
	}
}

func m2() {
	var arr = []int{1, 3, 5, 7, 9}
	for index, value := range arr {
		fmt.Println("index:", index, " value:", value)
	}
}

func m3() {
	arr_name := [5]*int{1: new(int), 3: new(int)}
	*arr_name[1] = 10
	*arr_name[3] = 30

	arr_name[4] = new(int)

	fmt.Println(*arr_name[1])
	fmt.Println(*arr_name[3])
	fmt.Println(*arr_name[4])
}

func m5() {
	var str_arr1 [3]*string
	str_arr2 := [3]*string{
		new(string),
		new(string),
		new(string),
	}

	*str_arr2[0] = "Perl"
	*str_arr2[1] = "Python"
	*str_arr2[2] = "Shell"

	str_arr1 = str_arr2
	fmt.Println(*str_arr1[1])

}

func m4() {
	values := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	values = append(values, row1)
	values = append(values, row2)

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[0][0])

	list := [2][2]string{}

	list[0][0] = "你好"
	list[0][1] = "欢迎"
	list[1][0] = "来到"
	list[1][0] = "杭州"

	fmt.Println(list)
}

func m6() {
	nums := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(nums[i][j])
		}
	}
}

func m7() {
	list := [][]string{}

	num1 := []string{"zhangs", "wang", "zhao"}
	num2 := []string{"li"}
	num3 := []string{"sum", "jin"}

	list = append(list, num1)
	list = append(list, num2)
	list = append(list, num3)

	for i := range list {
		fmt.Printf("list: %v\n", i)
		fmt.Println(list[i])
	}
}

func main() {
	// m1()
	// m2()
	// m3()
	// m4()
	// m6()
	m7()
}

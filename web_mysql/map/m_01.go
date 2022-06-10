package main

import (
	"fmt"
	"strconv"
	"strings"
)

func m1() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["name"] = "Marry"
	fmt.Println(a)
}

func m2() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "赣州"
	fmt.Println(cities)
}

func m3() {
	names := map[string]string{
		"name1": "小明",
		"name2": "小花",
		"name3": "小李",
	}
	fmt.Println(names)
}

func m4() {
	dict := map[string]string{"name": "zs", "address": "nj"}
	fmt.Println(dict)

	if len(dict) == 0 {
		fmt.Println("空")
	} else {
		fmt.Println("元素个数", len(dict))
	}
}

func m5() {
	Map1 := make(map[string]string)

	countryCapitalMap := make(map[string]string)

	Map1["France"] = "巴黎"
	Map1["Italy"] = "罗马"
	Map1["Japan"] = "东京"
	Map1["India"] = "新德里"
	Map1["Usa"] = "华盛顿"

	for key, value := range countryCapitalMap {
		if _, ok := Map1[key]; !ok {
			Map1[key] = value
		}
	}

	fmt.Println("===>", Map1)
}

func check() {
	colors := map[string]string{}
	colors["red"] = "红色"
	colors["blue"] = "蓝色"

	value, exe := colors["blue"]
	if exe {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}
}

func m6() {
	disc := map[string]int{"1": 10, "2": 20, "3": 30, "4": 40, "5": 50}
	value := disc["2"]
	if value != 0 {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}

	for k, v := range disc {
		fmt.Printf("key=%s,value=%d\n", k, v)
	}
	fmt.Println("-------------------------")
	//del
	delete(disc, "2")

	test(disc)

	for k, v := range disc {
		fmt.Printf("key=%s,value=%d\n", k, v)
	}
}

func m7() {
	students := make(map[string]map[string]string)

	students["01"] = make(map[string]string, 3)
	students["01"]["name"] = "tom"
	students["01"]["sex"] = "男"
	students["01"]["address"] = "rode1"

	students["02"] = make(map[string]string, 3)
	students["02"]["name"] = "marry"
	students["02"]["sex"] = "女"
	students["02"]["address"] = "rode2"

	fmt.Println(students)
	fmt.Println(students["02"]["name"])

}

func m8() {
	testmap := map[string][]string{}
	testmap["1"] = []string{"001", "002"}
	fmt.Println(testmap["1"])
}

func m9() {
	s := "how do you do"

	words := strings.Split(s, " ")

	a := make(map[string]int, 0)

	for _, word := range words {
		v, ok := a[word]
		if ok {
			a[word] = v + 1

		} else {
			a[word] = 1
		}
	}

	fmt.Println(a)
}

func m10() {
	var num int = 123456
	fmt.Printf("num 数据类型 %T \n", num)

	result := fmt.Sprintf("%d", num)
	fmt.Printf("类型 %T ", result)
}

func m11() {
	var str string = "123456"
	result, err := strconv.Atoi(str)
	if err == nil {
		fmt.Printf("数据 %T ", result)
	} else {
		fmt.Println(err)
	}
}

func test(m map[string]int) {
	m["2"] = 100
}

func main() {
	m1()
	m2()
	m3()
	m4()
	m5()
	check()
	m6()
	m7()
	m8()
	m9()
	m10()
	m11()
}

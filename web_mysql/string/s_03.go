package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	index()
	replace()
	count()
	repeat()
	toUpperOrLower()
	trim()
	split()
	itoa()
}

//位置索引
func index() {
	var (
		s   string = "adasfsddasscdsdkfjcdasdf"
		str string = "cd"
	)

	start := strings.Index(s, str)
	fmt.Println("位置：", start)
	end := strings.LastIndex(s, str)
	fmt.Println("end 位置：", end)
}

//替换
func replace() {
	var (
		s   string = "jajdjdjadcdjksdjcdsss"
		str string = "cd"
	)

	result := strings.Replace(s, str, "CD", -1)
	fmt.Println(result)
}

//统计次数
func count() {
	var (
		s   string = "jajdjdjadcdjksdjcdsss"
		str string = "cd"
	)

	count := strings.Count(s, str)
	fmt.Println("cd 次数：", count)
}

//重复
func repeat() {
	var (
		str string = "cd"
	)
	result := strings.Repeat(str, 3)
	fmt.Println(result)
}

//大小写
func toUpperOrLower() {
	var (
		str string = "abc"
	)

	result1 := strings.ToUpper(str)
	fmt.Println(result1)
	result2 := strings.ToLower(result1)
	fmt.Println(result2)

}

//去除字符
func trim() {
	var (
		s   string = "  abc"
		str string = "this is abc web"
	)

	fmt.Println("原长度：", len(s))
	//去除首部空白字符

	result := strings.TrimSpace(s)
	fmt.Println(result)
	fmt.Println("现长度：", len(result))

	//去除收尾和指定位置字符
	result1 := strings.Trim(str, "this")
	fmt.Println(result1)
	fmt.Println("现长度：", len(result1))

	//删除Left 和 Right
	result3 := strings.TrimLeft(str, "thisa ")
	fmt.Println(result3)
	result4 := strings.TrimRight(str, "_isweb ")
	fmt.Println(result4)
}

//切片
func split() {
	var s string = "abc:def:hig"
	result1 := strings.Split(s, ":")
	fmt.Println(" %T    %v ", result1, result1)

	str := strings.Join(result1, ",")
	fmt.Println("%T    %v ", str, str)
}

//数组处理
func itoa() {
	var num int = 123456
	fmt.Println("num 数据类型  %T \n", num)

	result := strconv.Itoa(num)
	fmt.Printf("result 数据类型 %T \n", result)
	fmt.Println(result)

	fmt.Println("---------------------------")
	var str string = "123456"
	result1, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println("数据类型 %T \n", result1)
	} else {
		fmt.Println(err)
	}
}

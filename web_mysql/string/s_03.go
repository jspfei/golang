package main

import (
	"fmt"
	"strings"
)

func main() {
	index()
	replace()
	count()
	repeat()
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

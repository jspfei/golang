package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	sleepTime()
	end := time.Now().UnixNano()
	fmt.Printf("cost: %d us", (end-start)/1000)
}
func sleepTime() {
	time.Sleep(time.Millisecond * 100)
}

func num2() {
	now := time.Now()
	times := now.Format("2006/1/02 15:04:05")
	fmt.Println(times)

	// 2022/4/05 14:45:30
	// 第一次按空格分割
	slice1 := strings.Fields(times)
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])
	// 第二次按 / 分割
	slice2 := strings.Split(slice1[0], "/")
	fmt.Println("年份:", slice2[0])
	fmt.Println("月份:", slice2[1])
	fmt.Println("日期:", slice2[2])
	fmt.Printf("slice2 数据类型 %T\n", slice2)

	// 奇数是会员日，偶数非会员日
	// 数据类型转换 string --> int
	// 判断是否有余数
	num, _ := strconv.Atoi(slice2[2])
	if num%2 != 0 {
		fmt.Println("会员日")
	} else {
		fmt.Println("非会员日")
	}

}

func num1() {
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 15:04:03"))
	fmt.Println(now.Format("2006/1/02 15:04"))
	fmt.Println(now.Format("2006/1/02"))

	times := now.Format("02/1/2006 15:04:05")
	fmt.Printf("times 数据类型 %T\n", times)
}

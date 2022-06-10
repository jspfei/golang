package main

import (
	"fmt"
)

func m1() {
	var (
		a  int = 33
		ip *int
	)

	ip = &a

	fmt.Println("a 的值 ：", a)
	fmt.Println("a 的地址 ：", &a)
	fmt.Println("指针 ip 的值 ：", ip)
	fmt.Println("指针 ip 指向的值 ：", *ip)

}

func m2() {
	a := []int{10, 100, 200}

	var ptr [3]*int

	for i := 0; i < 3; i++ {
		ptr[i] = &a[i]
	}

	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func m3() {
	var ptr [3]*int

	a := []int{10, 100, 200}

	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
		fmt.Printf("第 %d 个元素指针地址为 %d\n", i, &a[i])
	}

	for j := 0; j < len(ptr); j++ {
		fmt.Printf("a[%d] = %d\n", j, *ptr[j])
	}
}

func m4() {
	var a int = 100
	var b int = 200
	fmt.Println("原  a 值:", a)
	fmt.Println("原  b 值:", b)

	swap(&a, &b)
	fmt.Println("交换后 a 值;", a)
	fmt.Println("交换后 b 值;", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	m1()
	m2()
	m3()
	m4()
}

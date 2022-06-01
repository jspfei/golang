package main

import "fmt"

func main() {
	var (
		str1 = "hello"
		str2 = "world"
	)

	str3 := str1 + " " + str2
	fmt.Printf("数据 %s 长度 %d \n", str3, len(str3))
	fmt.Printf("数据 %s 类型 %T\n", str3)
	fmt.Println()

	tmp := []byte(str3)
	fmt.Printf("数据 %s 长度 %d\n", tmp, len(tmp))
	fmt.Printf("数据 %s 类型 %T\n", tmp)
	for i := 0; i < len(tmp); i++ {
		fmt.Println(tmp[i])
	}
	fmt.Println("字节数组转字符串：", string(tmp))
	fmt.Println()

	// 反转函数调用
	fmt.Println("反转：", Reverse(str3))

}

func Reverse(str string) string {
	var (
		result string
		leng   int = len(str)
	)
	for i := leng - 1; i >= 0; i-- {
		// Sprintf 转换数据类型为 string，使用 + 进行拼接
		// 重点：Sprintf 将任意类型转换成字符串
		result = result + fmt.Sprintf("%c", str[i])
	}
	return result

}

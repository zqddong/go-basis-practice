package main

import "fmt"

func main() {
	// 字符串可以通过 `+` 连接
	fmt.Println("go " + "lang")

	// 整数 浮点数
	fmt.Println("1+1", 1+1)
	fmt.Println("9.0/2.0 =", 9.0/2.0)

	// 布尔型
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

package main

import "fmt"

// 函数是 Go 的中心
func plus(a int, b int) int {
	// Go 需要明确返回 不会自动返回
	return a + b
}

func plus2(a, b, c int) int {
	return a + b + c
}

func main() {
	// 通过 `name(args)` 调用函数
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plus2(1, 2, 3)
	fmt.Println("1+2+3=", res)
}

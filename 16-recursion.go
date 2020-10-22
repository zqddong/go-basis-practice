package main

import "fmt"

// 递归 Recursion

// Go 支持递归 这是一个典型的阶乘示例

func fact(n int) int {
	fmt.Println("--", n)
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}

package main

import "fmt"

func main() {
	// var 声明一个或多个变量
	var a = "initial"
	fmt.Println(a)

	// 一次性声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go 自动推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	// 声明后没有给出对应的初始值 变量初始化为零值 一个 `int` 的零值是 0
	var e int
	fmt.Println(e)

	// `:=` 语法是声明并初始化变量的简写
	f := "short"
	fmt.Println(f)
}

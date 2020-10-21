package main

import "fmt"
import "math"

const s string = "this is constant"

//Go 支持字符、字符串、布尔和数值 常量 。
func main() {
	fmt.Println(s)
	// `const` 语句可以出现在任何 `var` 语句可以出现的地方
	const n = 500000000

	// 常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	// 数值型常量没有确定的类型，直到被给定，
	// 比如一次显式的类型转化。
	fmt.Println(int64(d))

	// 当上下文需要时 比如变量赋值或者函数调用
	// 一个数可以被给定一个类型
	fmt.Println(math.Sin(n))
}

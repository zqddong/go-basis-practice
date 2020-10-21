package main

import "fmt"

//在 Go 中，数组 是一个具有固定长度且编号的元素序列。
func main() {
	// 元素的长度和类型都是数组类型的一部分
	// 数组默认是零值的 对于 int 数组来说 也就是 `0`
	var a [5]int
	fmt.Println("emp:", a)

	// 使用 `array[index] = value` 语法设置数组
	// 指定的值 用 `array[index]` 得到值
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	// 使用内置函数 `len` 获取数组长度
	fmt.Println("array len:", len(a))

	// 使用语法在一行内声明并初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	// 数组类型是一维的 可以组合构造多维的数据结构

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}

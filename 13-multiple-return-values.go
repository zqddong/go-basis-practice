package main

import "fmt"

//多返回值函数

func values() (int, int) {
	return 3, 7
}

// Go 内建多返回值支持 用来同时返回一个函数的结果和错误信息
func main() {
	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

	// 仅需要返回值的一分部 使用空白标识符 `_` 忽略值
	_, c := values()
	fmt.Println(c)
}

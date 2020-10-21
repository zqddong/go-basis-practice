package main

import "fmt"

func main() {
	// 最基础的方式 单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的初始/条件/后续 `for` 循环
	for j := 10; j <= 16; j++ {
		fmt.Println(j)
	}

	// 不带条件的 `for` 循环一直重复执行 执行循环体出现 `break` 或 `return`
	for {
		fmt.Println("loop")
		break
	}

	// 使用 `continue` 跳出到下一个循环迭代
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

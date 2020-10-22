package main

import "fmt"

// 可变参数函数 调用时可以用任意数量的参数 fmt.Println 就是常见的变参函数

func sum(nums ...int) {
	fmt.Print(nums, ":")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// 含有多个值的 slice 把它们作为参数的调用方式 `func(slice...)`
	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}

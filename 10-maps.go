package main

import "fmt"

// map 是 Go 内置关联数据类型（在一些其他语言中称为哈希（hash）或字典（dict））
func main() {
	// 创建一个空的 map 需要使用内建的 `make`
	// `make(map[key-type]val-type)`
	m := make(map[string]int)

	// 使用典型的 `make[key] = val` 来设置键值对
	m["key1"] = 7
	m["key2"] = 13
	fmt.Println("map: ", m)

	// 获取
	v1 := m["key1"]
	fmt.Println("v1: ", v1)

	// 对一个 map 调用内置的 `len` 时，返回的是键值对的数目
	fmt.Println("len: ", len(m))

	// 内建函数 `delete` 可以从一个 map 中移除键值对
	delete(m, "key2")
	fmt.Println("delete map:", m)

	// 从一个 map 中取值时 可选第二返回值指示这个键是否在这个 map 中
	// 这可以用来消除键不存在和键有零值， 像 `0` 或 `""` 而产生的起义
	// 这里我们不要值，所以用 `_` 空白标识符（blank identifier）
	_, prs := m["key2"]
	fmt.Println("prs: ", prs)

	// 快速声明初始化一个新的 map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}

package main

import (
	"fmt"
	"go-basic-practice/php2go"
)

func main() {
	arr := map[interface{}]interface{}{"a": 1, "b": 2, "c": 3}
	keys := php2go.ArrayKeys(arr)
	fmt.Println("arr keys ", keys)

	valus := php2go.ArrayValues(arr)
	fmt.Println("arr values ", valus)
}

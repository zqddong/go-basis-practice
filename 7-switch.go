package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's weekday")
	}

	// 不带表达式的 `switch`  是实现 if/else 逻辑的另一种方式
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("now is am")
	default:
		fmt.Println("now is pm")
	}

	// 类型开关 `type/switch` 比较类型而非值 可以用来发现一个接口的类型值
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am a bool")
		case int:
			fmt.Println("i am a int")
		default:
			fmt.Printf(" don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
}

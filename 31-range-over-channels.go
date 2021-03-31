package main

import "fmt"

// 通道遍历
func main() {

    // 我们将遍历在 `queue` 通道中的两个值。
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // 这个 `range` 迭代从 `queue` 中得到的每个值。因为我们
    // 在前面 `close` 了这个通道，这个迭代会在接收完 2 个值
    // 之后结束。如果我们没有 `close` 它，我们将在这个循环中
    // 继续阻塞执行，等待接收第三个值
    for elem := range queue {
        fmt.Println(elem)
    }
}
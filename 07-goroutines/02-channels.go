package main

import "fmt"

// Channels
// 通道是一个类型化的管道，
// 可以通过它使用通道运算符 <- 发送和接收值。
/*
ch <- v   // 发送 v 到管道 ch 中
v := <-ch // 从管道 ch 中接收值，并赋值给 v
*/
// 跟 map 和 slice 一样，通道必须在使用之前创建
// ch := make(chan int)
// 默认，发送和接收是阻塞的, 直到对方准备好。
// 这允许 goroutines 在没有显式锁或条件变量的情况下进行同步。

// 示例代码将切片中的数字相加，在两个 goroutine 之间分配工作。
// 一旦两个 goroutine 都完成了计算，它就会计算最终结果。
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func testChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

// buffered channels
// 管道可以被缓存，提供缓存长度作为 make 的第二个参数来初始化一个带缓存的管道
func testBufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func main() {
	testChannel()	
	testBufferedChannel()
}
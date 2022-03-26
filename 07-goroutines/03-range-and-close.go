package main

import "fmt"

// range and close
// 发送者可以关闭通道以指示不再发送任何值。
// 接收者可以通过为接收表达式分配第二个参数来测试通道是否已关闭：之后
// v, ok := <-ch

// 循环 for i := range c 可以冲管道中重复的接收值直到管道被关闭

// 注意：只有发送者可以关闭一个管道，接收者不可以
// 注意：通常管道不需要关闭，只有在接受者需要被告知没有更多值的时候才需要关闭

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func testRangeAndCloseChanel() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	testRangeAndCloseChanel()
}
package main

import (
	"fmt"
	"time"
)

// select
// select 语句让 goroutine 等待多个通信操作。
// 一个 select 阻塞，直到它的一个 case 可以运行，然后它执行那个 case。
// 如果多个都准备好了，它会随机选择一个。
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func testSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

// default selection
// 在 select 中如果没有其它分支准备好，那么 default 分支会被执行
func testDefaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
		default:
			fmt.Println("                .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	testSelect()	
	testDefaultSelection()
}
package main

import (
	"fmt"
	"time"
)

// Goroutines
// goroutine 是由 Go 运行时管理的轻量级线程。
// go f(x, y, z)
// 开始一个新的 goroutine 运行
// f(x, y, z)
// 对于 x, y, z 运行在当前的 goroutine 中，而对于 f 则运行在新的 goroutine 中
// goroutine 运行在相同的地址空间，所以对共享内存的访问必须同步
// sync 包提供了有用的原语，尽管在 Go 中你不需要它们，因为还有其他原语。
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func testGoroutines() {
	go say("world")
	say("hello")
}

func main() {
	testGoroutines()	
}
package main

import "fmt"

// defer
// defer 语句将推迟到函数上下文 return 前执行

func testDefer() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}

// Stacking defers
// defer 函数调用被压入到一个栈中，当一个函数返回时，它的 defer 调用按后进后出的顺序执行。
func testStackDefers() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	testDefer()
	testStackDefers()
}
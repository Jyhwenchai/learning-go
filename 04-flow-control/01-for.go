package main

import "fmt"

// for 循环
// 基本的 for 循环由分号分隔的三个组件组成
// 1. 初始语句：在第一次迭代器执行
// 2. 条件表达式：每一次迭代时执行
// 3. post 语句：在每次迭代结束时执行
func testFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1	
	}
	fmt.Println(sum)
}

// for confinued
// 初始语句和 post 语句是可选的
func testForContinued() {
	sum := 0
	for ; sum < 1000; {
		sum += 1	
	}
	fmt.Println(sum)
}

// for 就是 go 中的 while 
// go 中没有 while 关键字，for 就是 while 循环
func testWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// 永远执行的 for 
func testForEver() {
	for {}
}

func main() {
	testFor()	
	testForContinued()
	testWhile()
}
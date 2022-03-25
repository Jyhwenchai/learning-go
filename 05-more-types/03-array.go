package main

import "fmt"

// array (数组)
// go 中数组类型用 [n]T 的方式表示，其中 T 为数组元素的类型，n 表示数组的长度
// 在 go 中，数组的长度是其类型的一部分，因此不能重新修改数组的长度
func testArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func main() {
	testArray()
}

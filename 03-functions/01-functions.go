package main

import "fmt"

// 函数以 func 关键字开始
// 示例中函数 add 带有两个 int 类型的参数，并返回 int 类型的值
func add(x int, y int) int {
	return x + y
}

// 共享类型
// 当函数有连续多个参数有相同的类型，那么可以使用下面的写法共享类型
func add2(x, y int) int {
	return x + y
}

// 多返回值
// 函数可以一次性返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
// 函数的返回值可以被命名，它们会被视为函数体开始定义的变量
func split(sum int) (x, y int) {
	x = sum *4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))	
	fmt.Println(swap("A", "B"))
	fmt.Println(split(10))
}

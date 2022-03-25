package main

import "fmt"

// pointer (指针)
// go 有指针，指针是保存内存地址的值的一个变量
// 与 C 不同，go 中的指针并不能做计算
func testPointer() {
	i, j := 42, 2701

	// go 中用 *T 来表示一个指针变量
	var p *int

	// 使用 & 运算符生成指向其操作数的指针
	p = &i

	// * 运算符表示指针的所指向的值。
	fmt.Println(*p)	// 解引用
	*p = 21	
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func main() {
	testPointer()	
}
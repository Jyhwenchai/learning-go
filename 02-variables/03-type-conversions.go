package main

import (
	"fmt"
)

// 类型表达式 T(v) 可以将值 v 的类型转换为类型 T
func testTypeConversation() {
	// var x, y int = 3, 4	
	// var f float64 = float64(x * y)
	// var z uint = uint(f)
	
	// 简便的写法
	x, y := 3, 4	
	f := float64(x * y)
	z := uint(f)
	fmt.Println(x, y, z)
}

// 类型引用
// 可以使用类型引用 := 为变量负责，变量可以根据右侧的值正确的推断出类型
func testTypeInference()  {
	v := 42
	fmt.Printf("v is of type %T", v)
}

// 常量
// go 中使用 const 来定义常量，常量不能使用 := 方式书写
const Pi = 3.14
func testConst() {
	const World = "World"
	fmt.Println("Hello, World")	
	fmt.Println("Happy", Pi, "Day")	
	
	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// 数值常量
// 数值类型常量可以表示高精度的值
// 一个无类型的常量根据其上下文推断出类型
const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func testNumericConstant() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func main() {
	testTypeConversation()
	testTypeInference()
	testConst()
	testNumericConstant()
}
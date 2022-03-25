package main

import (
	"fmt"
	"math"
)

// fucntion value
// 在 go 中函数也是值，它们可以像其他值一样被传递。
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func testFuncValue() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)	
	}	

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// function closure
// go 的函数可以使一个闭包，闭包是一个函数值，它从其主体外部引用变量。
// 该函数可以访问并分配给引用的变量；在这个意义上，函数是“绑定”到变量的。
func testFuncClosure() {
	adder := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x	// sum 被匿函数所持有生命周期延长
			return sum
		}
	}

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2 * i))
	}

}

// fibonacci closure
func testFibonacciClosure() {

	fibonacci := func() func() int {
		var current, next = 0, 1
		return func() int {
			ret := current
			current, next = next + current, ret
			return ret
		}
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
}
func main() {
	testFuncValue()
	testFuncClosure()
	testFibonacciClosure()	
}

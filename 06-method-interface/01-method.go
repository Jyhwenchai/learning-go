package main

import (
	"fmt"
	"math"
)

// method

// go 没有类，但是可以对类型添加方法
// 一个方法是一个有指定接收者的函数
// 接收者出现在 func 关键字和方法名称之间的自己的参数列表中。
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func testMethod() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// method are function
// 函数就是方法，只不过是拥有一个接受者的方法 
// 下面 ABS 用普通的函数实现上面方法的功能
func ABSFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func testMethodAreFunction() {
	v := Vertex{3, 4}
	fmt.Println(ABSFunc(v))
}

// 你也可以对非 struct 类型声明一个方法
type MyFloat float64
func (f MyFloat) ABS() float64 {
	if f < 0 {
		return float64(-f)
	}	 
	return float64(f)
}

func testMethodNotStruct() {
	f := MyFloat(-math.Sqrt2)	
	fmt.Println(f.ABS())
}

// pointer receiver
// 您可以使用指针接收器声明方法
// 这意味着接收者类型对于某些类型 T 具有语法 *T
// 带有指针接收器的方法可以修改接收器指向的值（就像 Scale 在这里所做的那样）。
// 由于方法经常需要修改它们的接收器，因此指针接收器比值接收器更常见。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func testMethodPointerReceiver() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// pointer and function
// 使用函数重写 Scale
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y *f
}

func testPointerAndFunction() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(ABSFunc(v))
}

// method and pointer indirection
// 比较之前的例子，对于函数相比较与方法必须多传递一个指针
// var v Vertex
// ScaleFunc(v, 5)	// error
// ScaleFunc(&v, 5)	// OK
//
// 具有指针接收器的方法在调用时将值或指针作为接收器：
// var v Vertex
// v.Scale(5) OK
// p := &v
// p.Scale(10) OK
func testMethodPointerIndirection() {
	v := Vertex{3, 4}
	v.Scale(2)
	Scale(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	Scale(p, 8)

	fmt.Println(v, p)
	
}

// method and pointer indirection 2
// 对于 p := &Vertex{3, 4}，它即可以作用于与函数也可以作用于方法
// 但是对于 v := Vertex{3, 4} 则不能进行下面的操作
// AbsFunc(v) OK
// AbsFunc(&v) error
//
// 另外对于带有指针接收者的方法，下面的操作也是有效的
// v := Vertex{3, 4}
// v.Abs() OK
// p := &v
// p.Abs()
func testMethodPointerIndirection2() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(ABSFunc(v))

	p := &Vertex{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(ABSFunc(*p))
}

// 选择值接收者还是指针接受者
// 使用指针接收器有两个原因。
// 第一个是让方法可以修改其接收器指向的值。
// 第二个是避免在每次方法调用时复制值。
// 例如，如果接收器是一个大型结构，这可能会更有效。
func main() {
	testMethod()	
	testMethodAreFunction()
	testMethodNotStruct()
	testMethodPointerReceiver()
	testPointerAndFunction()
	testMethodPointerIndirection()
	testMethodPointerIndirection2()
}
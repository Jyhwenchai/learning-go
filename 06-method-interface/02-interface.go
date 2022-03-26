package main

import (
	"fmt"
	"math"
)

// interface
// 一个借口类型可以定义一系列的方法签名
type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64
// 实现某个接口是隐式实现的，如下 Myfloat 实现了 Abs 方法，
// 该方法的签名与 Abser 中声明的完全相同，所以 Myfloat 就实现了该接口
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func testInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v 

	// a = v //error, 由于 Abs() 的接收者是 *Vertex 而 a 是一个值类型，所以不能赋值成功
	// 但是如果定义一个 Abs() 的接受者为 Vertex 类型，则 a = v 有效


	fmt.Println(a.Abs())
	fmt.Println(v.Abs())
}

// 在底层，接口值可以被认为是一个值和一个具体类型的元组：(value, type)
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func testInterfaceValue() {
	var i I
	describe := func(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

// interface values with nil underlying values
// 如果接口本身内部的具体值是 nil，则将使用 nil 接收器调用该方法。
type Z struct {
	S string
}

func (z *Z) M() {
	if z == nil {
		fmt.Println("<nil>")	
		return
	}
	fmt.Println(z.S)
}

func testNilReceiver() {

	describe := func(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}

	var i I

	var z *Z
	i = z
	describe(i)
	i.M()

	i = &Z{"hello"}
	describe(i)
	i.M()
}

// nil interface value
// 一个没有引用具体类型的 nil 的接口调用方法会引发运行时错误
func testInterfaceRuntimeError() {
	var i I
	i.M() // runtime error
}

func main() {
	testInterface()	
	testInterfaceValue()
	testNilReceiver()
	// testInterfaceRuntimeError()
}
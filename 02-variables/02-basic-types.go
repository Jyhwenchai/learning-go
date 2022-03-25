package main

import (
	"fmt"
	"math/cmplx"
)

// go 的基础数据类型有
/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16	uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名，用来表示 Unicode 码

float32 float64

complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt((-5 + 12i))
)


// 0 值
// 如果变量在声明的时候没有显示的给定一个值，那么默认会初始化为 0 值
/*
	numberic 类型的 0 值为：0
	boolean 类型的 0 值为：false
	string 类型的 0 值为：""
*/
var i int
var f float64
var b bool
var s string

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

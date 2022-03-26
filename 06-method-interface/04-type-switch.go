package main

import "fmt"

// type switches
// 类型开关是一种允许串联多个类型断言的结构。
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func testTypeSwitch() {
	do(21)
	do("hello")
	do(true)
}

func main() {
	testTypeSwitch()	
}
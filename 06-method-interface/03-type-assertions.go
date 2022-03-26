package main

import "fmt"

// type assertions
// 类型断言提供对接口值的底层具体值的访问
// t := i.(T)
// 该语句断言接口值 i 持有具体类型 T 并将底层 T 值分配给变量 t。
// 如果 i 没有持有 T，那么将触发一个 panic
// 该类型推断实际上会返回两个值，一个是推断成功时的类型，第二个是返回是否推断成功
// t, ok = i.(T)
// 这样如果 ok 为 false 是 t 为 T 类型的0值就不会 panic 了
func testTypeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
}

func main() {
	testTypeAssertion()	
}

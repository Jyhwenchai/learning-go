package main

import (
	"fmt"
	"strings"
)

// map
// map 是一种用 key 映射 value 的类型
// 其 0 值是 nil，nil map 没有 key，也不能添加 key
// 使用 make 方法可以创建一个 map
func testMap() {
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	m = make(map[string]Vertex, 0)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}

	fmt.Println(m["Bell Labs"])
}

// map literal
// map 字面量和 struct 很像，但是它的 key 是必须的
func testMapLiteral() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{ 
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google": {37.42202, -122.08408},	// 可以忽略类型名称 
	}

	fmt.Println(m)
}

// mutating map
// 插入或更新一个元素
// m[key] = elem
// 获取一个元素
// elem = m[key]
// 删除一个元素
// delete(m, key)
// 通过二值分配测试一个键是否存在：
// 如果 key 存在则 ok 为 true，否则为 false，此时 elem 为 map 元素类型的 0 值
// elem, ok = m[key]
func testMutatingMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// map exercise
func WordCount(s string) map[string]int {
	var wc = make(map[string]int)
	var strs = strings.Split(s, " ")

	for _, str := range strs {
		if _, ok := wc[str]; ok {
			wc[str] += 1
		} else {
			wc[str] = 1
		}
	}
	return wc
}


func CharacterCount(s string) map[string]int {
	var wc = make(map[string]int)
	for _, ch := range s {

		e := string(ch)
		if _, ok := wc[e]; ok {
			wc[e] += 1
		} else {
			wc[e] = 1
		}
	}
	return wc
}

func testMapExercise() {
	wc := WordCount("The quick brown fox jumped over the lazy dog.")
	cc := CharacterCount("I am learning Go!")

	fmt.Println(wc)
	fmt.Println(cc)
}
func main() {
	testMap()	
	testMapLiteral()
	testMutatingMap()
	testMapExercise()
}



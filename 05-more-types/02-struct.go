package main

import (
	"fmt"
)

// struct (结构体)
// go 中的结构体是一系列字段的集合

func testStruct() {
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})
}

// struct field (结构体字段)
// go 中使用 . 号来访问结构体的字段
func testStructField() {
	type Veritex struct {
		X int
		Y int
	}

	v := Veritex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

// 指向结构体的指针
// 结构体字段可以通过一个结构体指针进行访问
func testStructPointer() {
	type Veritex struct {
		X int
		Y int
	}

	v := Veritex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v.X)
}

// 结构体字面量
// 结构字面量通过列出其字段的值来表示新分配的结构体的值。
func testStructLiteral() {
	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1} // Y:0 是隐式的
		v3 = Vertex{}	// X:0 和 Y:0 都是隐式的
		p = &Vertex{1, 2}
	)

	fmt.Println(v1, p, v2, v3)
}

func main() {
	testStruct()
	testStructField()
	testStructPointer()
	testStructLiteral()
}

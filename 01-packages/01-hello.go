package main

import (
	"bufio"
	"fmt"
)

// 每一个 go 程序都是从 main 包中的 main 函数开始运行
func main() {
	fmt.Printf("Hello, World")
	var p *[]int = new([]int)      // allocates slice structure; *p == nil; rarely useful
	var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

	fmt.Println(p, v)

}

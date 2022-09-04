package main

import (
	"fmt"
	"strings"
)

// slice (切片)
// 数组的长度是固定的，与之不同的是切片拥有动态的长度
// go 中使用 []T 来表示切片类型
// 可以通过指定两个索引（下限和上限）形成切片，用冒号分隔：
// a[low : high]
// 在 low 到 high 的范围中，包含了 low 索引指向的值，但是不包含 high 所指向的值
func testSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array

	var s []int = primes[1:4]
	fmt.Println(s)
}

// 切片就像是一个数组的引用
// 切片不存储任何数据，它只是描述了底层数组的一个片段
// 改变切片的元素，其底层的数组的元素也会随之发生改变
// 多个切片可以共享一个底层的数组，一个切片中的元素发生改变，其它切片所引用到的元素也会发生改变
func testSliceRefenenceArray() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// slice literal (切片字面量)
// 切片字面量像是一个没有长度的数组字面量
// 数组字面量：[3]bool{true, false, false}
// 切片字面量：[]bool{true, false, false}
func testSliceLiteral() {
  q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// 值是结构体类型的切片
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}

// slice defaults (切片默认值)
// 使用切片时，你可以忽略它的上限和下限，默认的会使用默认值替代
// a[low:high] a[0:10] a[:10] a[0:] a[:]
func testSliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	i := s[:2]
	fmt.Println(i)

	j := s[1:]
	fmt.Println(j)

	k := s[:]
	fmt.Println(k)
}

// slice length and capacity (切片的长度和容量)
// 切片同时拥有长度和容量
// 长度: 表示当前切片拥有的元素数量
// 容量: 表示底层数组中元素的数量，从切片中的第一个元素开始计数
// 可以使用 len(s)、cap(s) 获取切片的长度和容量
// 可以通过重新切片来扩展切片的长度，前提是它有足够的容量
func testSliceLengthAndCapacity() {

	printSlice := func(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 使切片其长度为零。
	s = s[:0]
	printSlice(s)

	// 扩展切片长度
	s = s[:4]
	printSlice(s)

	// 删除切片前两个元素
	s = s[2:]
	printSlice(s)

}



// Nil slice (nil 切片)
// 切片的0值是 nil，一个 nil 的切片的长度和容量都是0，其没有引用一个底层的数组
func testNilSlice() {
	var s[]int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

// creating a slice with make
// 可以通过 make 方法来创建一个切片
// 通过这种方式创建切片，其底层会创建一个0值数组，然后一个切片会关联到这个数组
// a := make([]int, 5)
// 为了指定容量，可以传递第三个参数
// b := make([]int, 0, 5)
func testMakeSlice() {
	printSlice := func(s string, x []int) {
		fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
	}
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

// slices of slices (切片的切片)
// 切片可以包含任何类型，也包括切片
func testSliceOfSlice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][0] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// appending to a slice (切片追加元素)
// 可以通过 go 自带的函数为一个切片追加元素
// func append(s []T, vs ...T) []T
// 第一个参数是要追加元素的切片，vs 是可变参数，是切片元素的类型
func testSliceAppend() {
	printSlice := func(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
	
}
func main() {
	testSlice()
	testSliceRefenenceArray()
	testSliceLiteral()
	testSliceDefaults()
	testSliceLengthAndCapacity()
	testNilSlice()
	testMakeSlice()
	testSliceOfSlice()
	testSliceAppend()
}


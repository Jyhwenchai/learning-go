package main

import "fmt"

// go 中使用 var 定义变量，与函数参数列表相同其类型写在最后
// 变量可以定义在 package 或 function 级别
var c, python, java bool

// 变量可以在声明时进行初始化
var j, k int = 1, 2

// 如果在声明时进行了初始化，则可以省略类型；该变量将采用初始化程序的类型。
var js, perl, ts  = true, false, "typescript"

// 简短的变量声明
// 在函数内，可以使用 := 来替代 var 来声明一个隐式类型变量
func shortDeclaration() int {
	k := 3
	return k
}

// 在函数外，每个变量或函数的命名必须以 var， func 关键字开始， 简写 := 是不允许的 

func main() {
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(js, perl, ts)
}
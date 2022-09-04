package main

import (
	"fmt"
	"math"
)

// if
// if 语句与 for 类型，它不需要使用 () 包裹，但是需要 {}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func testIf() {
	fmt.Println(sqrt(2), sqrt(-4))
}

// if 是可以有一个简短的语句在其条件判断前
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func testIfShortStatement() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

// if 和 else
// if 后可以跟 else 语句
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func testIfElse() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}


func main() {
	testIf()
	testIfShortStatement()
	testIfElse()
}
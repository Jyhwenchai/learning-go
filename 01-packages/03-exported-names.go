// 在 go 中，如果一个变量的名称是以大写字母开头的，那么它是导出的，
// 例如 math 包中的 Pi

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
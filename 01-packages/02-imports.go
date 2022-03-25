package main

// 可以写多条的 import 语句导入需要的 package
// import "fmt"
// import "math"

// 但是更好多风格是使用下面的导入方式
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
package main

import "fmt"

// range
// range 可以用来迭代一个 slice 或 map
// range 进行迭代时，会返回两个值，第一个值是索引，第二个值是索引对应元素的拷贝
func testRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)	
	}
}

// 你可以使用 _ 符号来跳过索引或值
func testRangeConfinued() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(1) 	// == 2**i
	}

	// 忽略索引
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// 忽略值
	for i, _ := range pow {
		fmt.Printf("%d ", i)
	}

	fmt.Println()
	// 忽略值也可以这样简写 
	for i := range pow {
		fmt.Printf("%d ", i)
	}
}

// slice exercise
func testSliceExercese() {
	Pic := func (dx, dy int) [][]uint8 {
		var sdy = make([][]uint8, 0, dy)
		for j := 0; j < dy; j++ {
			var sdx = make([]uint8, 0, dx)
			for i := 0; i < dx; i++ {
				sdx = append(sdx, uint8(i+j))
			}
			sdy = append(sdy, sdx)
		}
		return sdy 
	}

	// run on https://go.dev/tour/moretypes/18
	// pic.show(Pic)	
	Pic(2, 5)
}

func main() {
	testRange()	
	testRangeConfinued()
}


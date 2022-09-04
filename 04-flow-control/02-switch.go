package main

import (
	"fmt"
	"runtime"
	"time"
)

// switch
// 和 if 语句类似，在 switch 语句中可以加入一条赋值语句

func testSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")	
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}

// 没有条件的 switch
func testNoConditionSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default: 
		fmt.Println("Good evening.")
	}
}
func main() {
	testSwitch()
	testNoConditionSwitch()
}
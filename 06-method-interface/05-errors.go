package main

import (
	"fmt"
	"time"
)

// errors
// Go 程序用 error 值表示错误状态。
// error 类型是一个内置接口：
/*
type error interface {
	Error() string
}
*/

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run () error {
	return &MyError{time.Now(), "it didn't work"}
}

func testError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	testError()	
}
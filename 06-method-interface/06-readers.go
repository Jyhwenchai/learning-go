package main

import (
	"fmt"
	"io"
	"strings"
)

// readers
// io 包指定 io.Reader 接口，代表数据流的读取端。
// 下面是 io.Reader 接口中的 Read 方法
// func (T) Read(b []byte) (n int, err error)

func testReader() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)

		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// exercise 
type MyReader struct{}
func (r MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 65
	}
	return len(bytes), nil
}

func main() {
	testReader()	
}
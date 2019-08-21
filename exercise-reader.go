package main

import (
	"golang.org/x/tour/reader"
	"strings"
)

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (m MyReader) Read(p []byte) (n int, err error) {
	r := strings.NewReader("A")
	return r.Read(p)
}

func main() {
	reader.Validate(MyReader{})
}

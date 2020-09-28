package base

import (
	"fmt"
	"testing"
)

/*
Go语言字符串的底层结构在reflect.StringHeader中定义：
string 不可变
type StringHeader struct {
    Data uintptr
    Len  int
}


*/
func Test(T *testing.T) {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]
	// 字符串虽然不是切片，但是支持切片操作
	// s1 := "hello, world"[:5]
	// s2 := "hello, world"[7:]
	fmt.Println(hello, world)
}

package test

import "testing"

type test struct{}

func test1() test {
	t1 := test{}
	println("t1", &t1)
	return t1
}

func test2() *test {
	t2 := test{}
	println("t2", &t2)
	return &t2
}

func TestMem(t *testing.T) {
	t1 := test1()
	t2 := test2()
	println("t1", &t1, "t2", &t2)
}

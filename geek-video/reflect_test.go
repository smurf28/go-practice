package test

import (
	"reflect"
	"testing"
)

type User struct {
	Name string
	Age  int32
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func TestCheckType(t *testing.T) {
	u := &User{"liuhuilin", 24}
	name := reflect.ValueOf(*u).FieldByName("Name")
	t.Log(name)
}

// 比较map 和 切片
func TestCompareEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 4: "three"}

	// t.Log(a == b)
	t.Log("a == a", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	t.Log("s1 == s2", reflect.DeepEqual(s1, s2))
}

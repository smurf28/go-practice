package test

import (
	"testing"
)

func TestMapWithFactor(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	t.Log(m[1](2), m[2](2))
}

func TestMapforSet(t *testing.T) {
	t.Log("args ...interface{}")
}

package test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// 单例模式, 一段代码确定只执行一次 懒汉式

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Println(unsafe.Pointer(obj))
			wg.Done()
		}()
	}
}

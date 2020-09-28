package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 使用场景: 多个任务用不同的方式处理同一请求,谁先完成就直接返回, (打车, 多种车一起叫)

func runTask(id int) string {
	time.Sleep(10 * time.Microsecond)
	return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	// 使用buff channel 来防止协程资源泄漏
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRes := ""
	for j := 0; j < numOfRunner; j++ {
		finalRes += <-ch + "\n"
	}
	return finalRes
}

func TestOneTask(t *testing.T) {
	t.Log("runtime :", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("end :", runtime.NumGoroutine())

}

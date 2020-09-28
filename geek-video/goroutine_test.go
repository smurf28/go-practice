package test

import (
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	counter := 0
	for i := 0; i < 500; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter is :%d", counter)
}

func TestThreadSafe(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter is :%d", counter)
}

// 多路选择select

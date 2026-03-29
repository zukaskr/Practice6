package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter1 int
	var mu sync.Mutex
	var wg1 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			mu.Lock()
			counter1++
			mu.Unlock()
		}()
	}
	wg1.Wait()
	fmt.Println("Mutex Counter Result:", counter1)

	var counter2 int64
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			atomic.AddInt64(&counter2, 1)
		}()
	}
	wg2.Wait()
	fmt.Println("Atomic Counter Result:", atomic.LoadInt64(&counter2))
}

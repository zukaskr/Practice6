package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Solution 1: sync.RWMutex")
	var mu sync.RWMutex
	unsafeMap := make(map[string]int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			mu.Lock()
			unsafeMap["key"] = key
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	mu.RLock()
	fmt.Printf("Value from RWMutex: %d\n", unsafeMap["key"])
	mu.RUnlock()

	fmt.Println("\nSolution 2: sync.Map")
	var sm sync.Map
	var wg2 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func(key int) {
			defer wg2.Done()
			sm.Store("key", key)
		}(i)
	}
	wg2.Wait()
	val, _ := sm.Load("key")
	fmt.Printf("Value from sync.Map: %v\n", val)
}

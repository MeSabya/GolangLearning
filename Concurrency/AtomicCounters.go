package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops int64

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 10; c++ {
				atomic.AddInt64(&ops, 1)

			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops", ops)

}

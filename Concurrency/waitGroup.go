package main

import (
	"fmt"
	"sync"
	"time"
)

/*
What is a WaitGroup? #
A WaitGroup blocks a program and waits for a set of goroutines to finish before
moving to the next steps of execution.

*/

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d Ended\n", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()

}

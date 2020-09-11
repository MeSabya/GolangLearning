//Fibonacci using channels and golang

package main

import "fmt"

func fibonacci(out chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case out <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quitting")
			return
		}
	}
}

func main() {
	out := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-out)
		}
		quit <- 0
	}()

	fibonacci(out, quit)

}

/*
Channels can also be used to send cancellation signals across goroutines so
that goroutines can stop their execution.

*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

var signal = make(chan bool)

func printNumbers() {
	for {
		select {
		case <-signal:
			return
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Default case executed ....")
		}
	}
}
func main() {
	go printNumbers()
	time.Sleep(time.Second)
	fmt.Println("Before Number of goroutines ", runtime.NumGoroutine())
	signal <- true
	fmt.Println("After Number of goroutines ", runtime.NumGoroutine())
	fmt.Println("Program Exited")
}

/*
When the main goroutine returns, we have only 1 active goroutine. This way,
we can clean up any resources hogged by the printNumbers goroutine. This way
of communication between two goroutines can be used for simple operations like the above.

However, when there is a tight relationship between running goroutines and
one goroutine can start other goroutines,
then it becomes tremulously difficult to use channels to send cancellation signals.

if the printNumbers goroutine starts another goroutine, the main goroutine
won’t have any idea about it. So it becomes the responsibility of the PrintNumbers
goroutine to send a cancellation signal to it.
This gets more complicated when that goroutine spawns other goroutines.

This is where context package comes in.
The sole purpose of the context package is to carry out the cancellation
signal across goroutines no matter how they were spawned, context got them covered.
*/

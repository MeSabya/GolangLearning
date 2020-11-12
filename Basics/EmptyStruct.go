/*
An empty struct is a struct type without fields struct{}.
The cool thing about an empty structure is that it occupies zero bytes of storage.

The short version is:
The size of a struct is the sum of the size of the types of its fields,
since there are no fields: no size!

"Every time you have channels used to signal something, rather than exchanging values,
you can use empty structs.
This will also make your code clearer,
explicitly showing that you do not care about the values IN the channel."

*/

/*
Usually when people create semaphores in GO, you see something like this
sem := make(chan bool, numberOfSemaphores) or sem := make(chan int, numberOfSemaphores)
Then you can do
sem <- true or sem <- 1.
This is fine.
But what you can also do is:
sem := make(chan struct{}, numberOfSemaphores) and then
sem <- struct{}{}.

Basically you are declaring an array of empty structs, which occupies no storage.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const halfSecond = 500 * time.Millisecond

func main() {
	token := make(chan struct{}, 10)
	var wg sync.WaitGroup
	wg.Add(50)

	for i := 0; i < 50; i++ {
		go func() {
			token <- struct{}{}
			defer wg.Done()
			//fmt.Println("i value is", i)
			time.Sleep(halfSecond)
			<-token
		}()

		fmt.Println("Number of goroutines running", runtime.NumGoroutine())
	}

	start := time.Now()
	wg.Wait()
	fmt.Printf("It took me %d\n", time.Since(start))
}

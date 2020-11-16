/*
An object of interface type context.Context is used to carry deadlines,
cancelation signals, and values between concurrently running processes.
This object provides a Done signal channel that closes when we call the
cancel function received during the creation of this object.
This channel is be retrieved using the Done() method of the Context object.
*/

/*
// A Context carries a deadline, cancelation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
*/

/*
context.Background() if you only need it run as it is
context.WithValue(...) if you need to give some value to the process
context.WithCancel(...) if you want to able to cancel the process
context.WithDeadline(...) if you want the process done at some deadline time.
context.WithTimeout(...) if you want the process done after some duration(some with WithDeadline )
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	finish := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		finish <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Done")
	case <-finish:
		fmt.Println("Finished")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		time.Sleep(10 * time.Microsecond)
		cancel()
	}()

	process(ctx)
	if ctx.Err() != nil {
		fmt.Println(ctx.Err().Error())
	}

}

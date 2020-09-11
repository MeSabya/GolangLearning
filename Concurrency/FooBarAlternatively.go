//Print foo bar alternatively
package main

import (
	"fmt"
)

type FooBar struct {
	n            int
	streamFooBar chan struct{}
	streamBarFoo chan struct{}
	streamToEnd  chan struct{}
}

func (fooBar *FooBar) printFoo(print func()) {

	for i := 0; i < fooBar.n; i++ {
		<-fooBar.streamBarFoo
		print()
		fooBar.streamFooBar <- struct{}{}
	}
	<-fooBar.streamBarFoo
}

func (fooBar *FooBar) printBar(print func()) {
	for i := 0; i < fooBar.n; i++ {
		<-fooBar.streamFooBar
		print()
		fooBar.streamBarFoo <- struct{}{}
	}

	fooBar.streamToEnd <- struct{}{}
}

func main() {

	var PrintFooBar = func(times int) {
		fooBar := &FooBar{
			n:            times,
			streamFooBar: make(chan struct{}),
			streamBarFoo: make(chan struct{}),
			streamToEnd:  make(chan struct{}),
		}
		go fooBar.printFoo(func() { fmt.Printf("Foo") })
		go fooBar.printBar(func() { fmt.Printf("Bar") })
		//What is struct{}{} ?
		//The struct{} is a struct type with zero elements.
		//struct{}{} on the other hand is a composite literal.
		//it constructs a value of type struct{}. A composite literal
		//constructs values for types such as structs, arrays, maps and slices
		fooBar.streamBarFoo <- struct{}{}
		<-fooBar.streamToEnd
		fmt.Println()

	}

	testCase := []int{0, 1, 2, 3, 4, 5, 7}

	for _, repeat := range testCase {
		fmt.Printf("Repeat %d: ", repeat)
		PrintFooBar(repeat)
	}

}

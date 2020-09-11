package main

import "fmt"

func test2() {
	fmt.Println("test2() function activated")
}

//It is returning a function
func test() func() {
	return test2
}

//It is returning a lamda function
func test1() func() int {
	fmt.Println("Returning test2()")
	return func() int {
		return 42
	}
}

func main() {
	fn := test1()
	//Call the function
	fmt.Println(fn())
}

package main

import "fmt"

//https://www.calhoun.io/how-to-use-slice-capacity-and-length-in-go/
/*
* This program has three variation to test .
* var vals []int with capacity and lenght is 0
* vals := make([]int, 5), both len and capacity is 5
* vals := make([]int, 0, 5) len = 0, capacity is 5
* check the len and capacity and vals
 */
func main() {
	//var vals []int
	//vals := make([]int, 5)
	vals := make([]int, 0, 5)
	//fmt.Println("The length of our slice is:", len(vals))
	//fmt.Println("The capacity of our slice is:", cap(vals))
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
		//fmt.Println("Printing the vals:", vals)
		//fmt.Println("The length of our slice is:", len(vals))
		//fmt.Println("The capacity of our slice is:", cap(vals))

	}

	// Add a new item to our array
	vals = append(vals, 123)
	fmt.Println("The length of our slice is:", len(vals))
	fmt.Println("The capacity of our slice is:", cap(vals))

	// Accessing items is the same as an array
	fmt.Println(vals[5])
	fmt.Println(vals[2])

}

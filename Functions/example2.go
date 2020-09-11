//Write a function foo which takes []int as variable number of arguments
//Write another function which will take []int as it is , returns the sum of ele
//ments in both the cases

package main

import "fmt"

func foo(xi ...int) int {
	sum := 0
	for _, val := range xi {
		sum += val
	}

	return sum
}
func main() {
	val := []int{2, 3, 4}
	sum := foo(val...) //This is the way to pass variable arguments
	fmt.Println(sum)

}

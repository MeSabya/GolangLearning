package main

import (
	"fmt"
)

var a = make([]int, 7, 9)

// A slice is a descriptor of an array segment.
// It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
// The length is the number of elements referred to by the slice.
// The capacity is the number of elements in the underlying array (beginning at the element referred to by the slice pointer).
// |-> Refer to: https://blog.golang.org/go-slices-usage-and-internals -> "Slice internals" section

func Test(slice []int) {
	// slice receives a copy of slice `a` which point to the same array as slice `a`
	slice[6] = 10
	slice = append(slice, 100)
	// since `slice` capacity is 8 & length is 7, it can add 100 and make the length 8
	fmt.Println(slice, len(slice), cap(slice), " << Test 1")
	slice = append(slice, 200)
	// since `slice` capacity is 8 & length also 8, slice has to make a new slice
	// - with double of size with point to new array (see Reference 1 below).
	// (I'm also confused, why not (n+1)*2=20). But make a new slice of 16 capacity).
	slice[6] = 13 // make sure, it's a new slice :)
	fmt.Println(slice, len(slice), cap(slice), " << Test 2")
}

func main() {
	for i := 0; i < 7; i++ {
		a[i] = i
	}

	fmt.Println(a, len(a), cap(a))
	Test(a)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a[:cap(a)], len(a), cap(a))
	// fmt.Println(a[:cap(a)+1], len(a), cap(a)) -> this'll not work
}

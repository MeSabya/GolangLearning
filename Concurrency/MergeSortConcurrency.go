package main

import (
	"fmt"
	"time"
)

func merge_(left []int, right []int) (result []int) {
	//Here left and right arrays are already sorted
	//We need to make a sorted array out of the two sorted arrays
	left_len := len(left)
	right_len := len(right)
	i, j, k := 0, 0, 0
	result = make([]int, left_len+right_len)

	for i < left_len && j < right_len {
		if left[i] <= right[j] {
			result[k] = left[i]
			k = k + 1
			i = i + 1
		} else {
			result[k] = right[j]
			k = k + 1
			j = j + 1

		}
	}

	for i < left_len {
		result[k] = left[i]
		k = k + 1
		i = i + 1
	}
	for j < right_len {
		result[k] = right[j]
		k = k + 1
		j = j + 1
	}

	return
}

/*
Just look at the function above, it looks like it is implemented in C or C++,
It can be improved more in a golang way. The improved version is:
*/

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		} else if len(right) == 0 {
			return append(result, left...)
		} else if left[0] < right[0] {
			result = append(result, left[0])
			left = left[0:]
		} else {
			result = append(result, right[0])
			right = right[0:]
		}
	}

	return result
}

func singleMergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	return merge(singleMergeSort(arr[:mid]), singleMergeSort(arr[mid:]))
}

func concurrentMerge(arr []int, result chan []int) {
	if len(arr) < 2 {
		result <- arr
		return
	}

	left := make(chan []int)
	right := make(chan []int)
	middle := len(arr) / 2

	go concurrentMerge(arr[:middle], left)
	go concurrentMerge(arr[middle:], right)

	ldata := <-left
	rdata := <-right

	close(left)
	close(right)

	result <- merge(ldata, rdata)
	return
}

func runConcurrentMerge(arr []int) []int {
	result := make(chan []int)
	go concurrentMerge(arr, result)
	res := <-result

	return res

}

func main() {
	arr := []int{3, 1, 4, 6, 7}
	start := time.Now()
	//result := singleMergeSort(arr)
	result := runConcurrentMerge(arr)
	fmt.Println("Total time taken ", time.Since(start))
	fmt.Println(result)

}

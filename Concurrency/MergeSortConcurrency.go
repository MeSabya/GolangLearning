package main

import (
	"fmt"
	"time"
)

func merge(left []int, right []int) (result []int) {
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

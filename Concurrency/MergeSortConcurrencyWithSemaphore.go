/*
In MergeSortConcurrency we have created so many go routines.package Concurrency
In this program using semaphore we can control the spawning of multiple
goroutines.
*/
package main

import (
	"fmt"
	"sync"
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

func runConcurrentMergeWithSem(arr []int, sem chan struct{}) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	wg := sync.WaitGroup{}
	wg.Add(2)

	var ldata []int
	var rdata []int

	select {
	case sem <- struct{}{}:
		go func() {
			ldata = runConcurrentMergeWithSem(arr[:mid], sem)
			<-sem
			wg.Done()
		}()
	default:
		ldata = singleMergeSort(arr[:mid])
		wg.Done()

	}

	select {
	case sem <- struct{}{}:
		go func() {
			rdata = runConcurrentMergeWithSem(arr[mid:], sem)
			<-sem
			wg.Done()
		}()
	default:
		rdata = singleMergeSort(arr[mid:])
		wg.Done()

	}

	wg.Wait()
	return merge(ldata, rdata)
}

func runConcurrentMerge(arr []int) []int {
	//Usage of creating a channel using an empty struct is discussed in
	//https://github.com/MeSabya/GolangLearning/blob/master/Basics/EmptyStruct.go
	sem := make(chan struct{}, 4) //At once we want to spawn maximum 4 goroutines
	return runConcurrentMergeWithSem(arr, sem)
}
func main() {
	arr := []int{3, 1, 4, 6, 7}
	start := time.Now()
	//result := singleMergeSort(arr)
	result := runConcurrentMerge(arr)
	fmt.Println("Total time taken ", time.Since(start))
	fmt.Println(result)

}

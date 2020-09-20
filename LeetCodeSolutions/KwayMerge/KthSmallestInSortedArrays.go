/*Find Kth smallest element in a list of sorted arrays ,
we have discussed priority queue for list of linked lists, here we will discuss about
priority queue over a list of sorted arrays. The basic difference between these two solutions
are in case of linked lists , we can reach to the next node value , but in case of
array , we need to maintain the next element in a data structure.
*/

/*
Simillar Problems:
Problem 1: Given ‘M’ sorted arrays, find the median number among all arrays.

Solution: This problem is similar to our parent problem with K=Median.
So if there are ‘N’ total numbers in all the arrays we need to find the K’th
minimum number where K=N/2K=N/2.

Problem 2: Given a list of ‘K’ sorted arrays, merge them into one sorted list.

Solution: This problem is similar to Merge K Sorted Lists except that the input is
a list of arrays compared to LinkedLists. To handle this, we can use a similar
approach as discussed in our parent problem by keeping a track of the array
and the element indices.

Problem3 :
Kth Smallest Number in a Sorted Matrix
Given an N * NN∗N matrix where each row and column is sorted in ascending order, find the Kth smallest element in the matrix.

Example 1:

Input: Matrix=[
    [2, 6, 8],
    [3, 7, 10],
    [5, 8, 11]
  ],
  K=5
Output: 7
Explanation: The 5th smallest number in the matrix is 7.

*/

package main

import (
	"container/heap"
	"fmt"
)

type ArrayContainer struct {
	Element int
	Index   int
	List    []int
}

type PriorityQueue []*ArrayContainer

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Element < pq[j].Element
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

//now heap overidden methods
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	lastval := old[len(*pq)-1]
	*pq = old[:len(*pq)-1]
	return lastval
}

func (pq *PriorityQueue) Push(val interface{}) {
	newval := val.(*ArrayContainer)
	*pq = append(*pq, newval)
}

func findKthSmallest(arrs [][]int, k int) int {
	pq := make(PriorityQueue, 0)

	for idx := 0; idx < len(arrs); idx++ {
		pq = append(pq, &ArrayContainer{arrs[idx][0], 0, arrs[idx]})
	}

	if (len(pq)) == 0 {
		return -1
	}
	heap.Init(&pq)
	count := 0
	for len(pq) > 0 {
		array := heap.Pop(&pq).(*ArrayContainer)
		ele, idx, list := array.Element, array.Index, array.List //TODO check for unpack operator
		count++
		if count == k-1 {
			return ele
		}

		//Here process the next element from the list.
		if len(list) > idx+1 {
			heap.Push(&pq, &ArrayContainer{list[idx+1], idx + 1, list})
		}

	}

	return -1

}

func main() {
	arr1 := []int{2, 6, 8}
	arr2 := []int{3, 6, 7}
	arr3 := []int{1, 3, 4}

	//var listOfArr [][]int
	fmt.Println(findKthSmallest([][]int{arr1, arr2, arr3}, 5))
}

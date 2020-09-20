/*Find Kth smallest element in a list of sorted arrays ,
we have discussed priority queue for list of linked lists, here we will discuss about
priority queue over a list of sorted arrays. The basic difference between these two solutions
are in case of linked lists , we can reach to the next node value , but in case of
array , we need to maintain the next element in a data structure.
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
	lastval := *pq[len(old)-1]
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

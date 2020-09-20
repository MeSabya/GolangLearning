package main

import (
	"container/heap"
	"fmt"
)

//Merge K sorted lists
//What should be the aproach?
//Push every list to min heap, (So figureout how to implement a minheap in golang)
//Until heap is empty, need to pop every element from heap and push next possible element.

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

/*
As discussed in the PriorityQueue.go file , we need to override 3methods to implement
sorting interface, those are : Len(), Less(i, j int), Swap(i, j int)
We need to override Push() and Pop() methods to implement heap interface.
*/
//Lets start defining the methods , starting with sort interface
func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

//Lets define overroide methods of heap interface
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	lastnode := old[len(*pq)-1]
	*pq = old[:len(*pq)-1]
	return lastnode
}

func (pq *PriorityQueue) Push(newnode interface{}) {
	node := newnode.(*ListNode)
	*pq = append(*pq, node)
}
func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	for _, node := range lists {
		if node != nil {
			pq = append(pq, node)
		}
	}

	if len(pq) == 0 {
		return nil
	}

	heap.Init(&pq)
	head := &ListNode{}
	dummyhead := head

	for len(pq) > 0 {
		minnode := heap.Pop(&pq).(*ListNode)
		head.Next = minnode
		head = head.Next

		if minnode.Next != nil {
			heap.Push(&pq, minnode.Next)
		}

	}

	return dummyhead.Next
}

func main() {
	node3 := &ListNode{6, nil}
	node2 := &ListNode{4, node3}
	node1 := &ListNode{1, node2}

	node5 := &ListNode{5, nil}
	node4 := &ListNode{2, node5}

	nodes := []*ListNode{node1, node4}
	mergedlists := mergeKLists(nodes)
	if mergedlists != nil {
		temp := mergedlists
		for temp != nil {
			fmt.Printf("%d", temp.Val)
			temp = temp.Next
		}
	}

}

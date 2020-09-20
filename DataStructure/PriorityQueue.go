//Priority Queue in Golang

/*
In the Push and Pop method we are using interface
Learn these :
interface{} is the empty interface type

[]interface{} is a slice of type empty interface

interface{}{} is an empty interface type composite literal

[]interface{}{} is a slice of type empty interface composite literals

What does interface{} meaning in Push and Pop operations ??
interface{} means you can put value of any type, including your own custom type. All types in Go satisfy an empty interface (interface{} is an empty interface).
In your example, Msg field can have value of any type.

Example:
package main

import (
    "fmt"
)

type Body struct {
    Msg interface{}
}

func main() {
    b := Body{}
    b.Msg = "5"
    fmt.Printf("%#v %T \n", b.Msg, b.Msg) // Output: "5" string
    b.Msg = 5
    fmt.Printf("%#v %T", b.Msg, b.Msg) //Output:  5 int
}
*/

package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Name   string
	Expiry int
	Index  int
}

type PriorityQueue []*Item

//In order to sort the priority queue , implement the
/* type Interface interface {
   // Len is the number of elements in the collection.
   Len() int
   // Less reports whether the element with
   // index i should sort before the element with index j.
   Less(i, j int) bool
   // Swap swaps the elements with indexes i and j.
   Swap(i, j int)
*/

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Expiry < pq[j].Expiry
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]

}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	//n := len(*pq)
	item := x.(*Item)
	*pq = append(*pq, item)
}

func main() {
	listItems := []*Item{
		{Name: "Carrot", Expiry: 30},
		{Name: "Potato", Expiry: 45},
		{Name: "Rice", Expiry: 100},
		{Name: "Spinach", Expiry: 5},
	}

	priorityQueue := make(PriorityQueue, len(listItems))

	for i, item := range listItems {
		priorityQueue[i] = item
	}
	/*
	* Here couple of things need to be considered :
	* heap works on pointers , for example: Both heap push and pop works on pointers
	* See the signature.
	* We should know how to work on interface. For example heap pop returns an interface
	* It should be converted to corresponding type of object.
	* heap.Pop(&priorityQueue).(*Item) , here Item is a pointer type.
	* Because while inserting into priorityQueue, we are inserting a pointer of Item type.
	 */
	heap.Init(&priorityQueue)

	for priorityQueue.Len() > 0 {
		item := heap.Pop(&priorityQueue).(*Item)
		fmt.Printf("Name %s Expiry:%d\n", item.Name, item.Expiry)
	}

}

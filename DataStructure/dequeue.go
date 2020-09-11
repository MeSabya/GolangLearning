package dequeue

import (
	"container/list"
)

type Deque struct {
	dequeue  *list.List
	capacity int
}

func NewDequeue() *Deque {
	return NewCappedDequeue(-1)
}

func NewCappedDequeue(capacity int) *Deque {
	return &Deque{
		dequeue:  list.New(),
		capacity: capacity,
	}

}

//Append inserts elements at the back of the Dequeue in a O(1) time complexity
//returns true if successful otherwise returns false
func (d *Deque) Append(item interface{}) bool {
	if d.capacity < 0 || d.dequeue.Len() < d.capacity {
		d.dequeue.PushBack(item)
		return true
	}

	return false
}

/*
In this Program we can learn
1. how to sort a 2d slice in place.
2. How to add more conditions in Less() function.
*/

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type journeyInfo struct {
	paxCount int
	source   int
	dest     int
}

type journeys []journeyInfo

func (j journeys) Len() int {
	return len(j)
}
func (jo journeys) Less(i, j int) bool {
	return jo[i].source < jo[j].source
}

func (jo journeys) Swap(i, j int) {
	jo[i], jo[j] = jo[j], jo[i]
}

type Item struct {
	dest int
	pax  int
}
type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (jo PriorityQueue) Less(i, j int) bool {
	if jo[i].dest < jo[j].dest {
		return true
	} else if jo[i].dest == jo[j].dest {
		return jo[i].pax > jo[j].pax
	}

	return false
}

func (jo PriorityQueue) Swap(i, j int) {
	jo[i], jo[j] = jo[j], jo[i]
}

func (jo *PriorityQueue) Push(item interface{}) {
	x := item.(Item)
	*jo = append(*jo, x)
}

func (jo *PriorityQueue) Pop() interface{} {
	old := *jo
	n := len(old)
	x := old[n-1]
	*jo = old[0 : n-1]

	return x
}

func carPooling(trips [][]int, capacity int) bool {
	/*journey := make(journeys,0)
	  for _, trip := range trips{
	      journey = append(journey, journeyInfo{trip[0], trip[1], trip[2]})
	  }
	  sort.Sort(journey)*/
	sort.Slice(trips, func(i, j int) bool {
		if trips[i][1] == trips[j][1] {
			return trips[i][2] < trips[j][2]
		}
		return trips[i][1] < trips[j][1]
	})
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, trip := range trips {
		fmt.Println("trip is", trip[0], trip[1], trip[2])
		if len(pq) > 0 {
			//fmt.Println("pq[0] ele are", pq[0].dest, pq[0].pax)
		}
		for len(pq) > 0 && pq[0].dest <= trip[1] {
			//fmt.Println("First pq.dest is", pq[0].dest)
			jinfo := heap.Pop(&pq).(Item)
			fmt.Println("Poped element are", jinfo.dest, jinfo.pax)
			capacity = capacity + jinfo.pax
		}

		//fmt.Println("Capacity is", capacity)
		if trip[0] <= capacity {
			capacity = capacity - trip[0]
			fmt.Println("Pushed into pq", trip[2], trip[0])
			heap.Push(&pq, Item{trip[2], trip[0]})
		} else {
			return false
		}
	}

	return true

}

func main() {

	/*ret := carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4)
	if ret != false {
		fmt.Println("Testcase1 failed")
	}
	ret = carPooling([][]int{{3, 2, 7},
		{3, 7, 9},
		{8, 3, 9},
	}, 11)
	if ret != true {
		fmt.Println("Testcase2 failed")
	}*/

	//This is the most important TC , where departure time is same.
	//In that case it should process the highest pax trip first.
	//So here between {7,5,6}, {10, 1, 6}
	tc3 := [][]int{{7, 5, 6}, {6, 7, 8}, {10, 1, 6}}
	ret1 := carPooling(tc3, 16)
	if ret1 != false {
		fmt.Println("TC3 Failed")
	}

	//Probable TCs are:
	//Same start , diff dest,
	//diff start, same dest

}

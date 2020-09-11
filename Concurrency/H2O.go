package main

import (
	"fmt"
	"sync"
)

type H2O struct {
	hCount     int
	oCount     int
	count      int
	cond       *sync.Cond
	exitSignal chan struct{}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	for i := 0; i < h2o.hCount; i++ {
		h2o.cond.L.Lock()
		for h2o.count >= 2 {
			h2o.cond.Wait()
		}

		releaseHydrogen()
		h2o.count++
		h2o.cond.L.Unlock()
		h2o.cond.Broadcast()
	}
	h2o.cond.Signal()
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	for i := 0; i < h2o.oCount; i++ {
		h2o.cond.L.Lock()
		for h2o.count <= 0 {
			h2o.cond.Wait()
		}

		releaseOxygen()
		h2o.count -= 2
		h2o.cond.L.Unlock()
		h2o.cond.Broadcast()
	}
	h2o.exitSignal <- struct{}{}
}

func main() {
	barrier := func(hCount, oCount int) {
		h2o := H2O{
			hCount:     hCount,
			oCount:     oCount,
			count:      0,
			cond:       sync.NewCond(&sync.Mutex{}),
			exitSignal: make(chan struct{}),
		}
		go h2o.hydrogen(func() { fmt.Println("H") })
		go h2o.oxygen(func() { fmt.Println("O") })
		<-h2o.exitSignal

	}

	/*type pair struct {
		Vals [2]interface{}
	}

	nums := []pair{pair{[2]interface{}{2, 1}}, pair{[2]interface{}{4, 2}},
		pair{[2]interface{}{9, 3}}}
	*/

	/*
	* This is the way to implement tuple in golang
	 */
	type pair [2]interface{}
	nums := []pair{pair{2, 1}, pair{4, 2}, pair{9, 3}}

	for _, pairO := range nums {
		//fmt.Println("Hydrogen and Oxygen count is ", pairO.Vals[0].(int), pairO.Vals[1].(int))

		//barrier(pairO.Vals[0].(int), pairO.Vals[1].(int))

		// Here we convert from interface{} to int type,
		//Otherwise we will get some error
		barrier(pairO[0].(int), pairO[1].(int))

	}

}

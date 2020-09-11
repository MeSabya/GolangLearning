package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := sync.Mutex{}
	cond := sync.NewCond(&m)

	//func1
	go func() {
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait()
		}
		fmt.Println(sharedRsc["SABYA"])
		cond.L.Unlock()
		wg.Done()

	}()

	//func2
	go func() {
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait()
		}
		fmt.Println(sharedRsc["GUDUL"])
		cond.L.Unlock()
		wg.Done()

	}()

	cond.L.Lock()
	sharedRsc["SABYA"] = 31
	sharedRsc["GUDUL"] = 25
	cond.Broadcast()
	cond.L.Unlock()
	wg.Wait()

}

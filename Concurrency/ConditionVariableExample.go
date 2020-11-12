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
		fmt.Println("fun1")
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait()
		}
		fmt.Println("fun1::", sharedRsc["SABYA"])
		cond.L.Unlock()
		wg.Done()

	}()

	//func2
	go func() {
		fmt.Println("fun2")
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait()
		}
		fmt.Println("Fun2::", sharedRsc["GUDUL"])
		cond.L.Unlock()
		wg.Done()

	}()

	fmt.Println("Before lock")
	cond.L.Lock()
	fmt.Println("After lock")
	sharedRsc["SABYA"] = 31
	sharedRsc["GUDUL"] = 25
	cond.Broadcast()
	cond.L.Unlock()
	wg.Wait()

}

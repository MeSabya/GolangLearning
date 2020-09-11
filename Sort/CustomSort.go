//How to sort Person structure

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	age  int
	name string
}

type nameSortedPs []Person

func (arr nameSortedPs) Len() int {
	return len(arr)
}

func (arr nameSortedPs) Less(i, j int) bool {
	//irune, _ := utf8.DecodeRuneInString(arr[i].name)
	//jrune, _ := utf8.DecodeRuneInString(arr[j].name)

	//return int32(irune) < int32(jrune)
	return arr[i].name[0] < arr[j].name[0]

}

func (arr nameSortedPs) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	persons := []Person{
		{name: "sabya", age: 30}, {name: "Gudul", age: 25},
	}

	sort.Sort(nameSortedPs(persons))

	fmt.Println(persons)
}

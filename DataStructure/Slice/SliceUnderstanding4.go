package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func create(iterations int) []int {
	a := make([]int, 0)
	for i := 0; i < iterations; i++ {
		a = append(a, i)
	}
	return a
}

func main() {
	sliceFromArray()
	//sliceFromLiteral()

}

func sliceFromArray() {
	//fmt.Printf("** NOT working as expected: **\n\n")
	// this doesnt work
	i := create(11)
	hdri := (*reflect.SliceHeader)(unsafe.Pointer(&i))
	data := *(*[4]int)(unsafe.Pointer(hdri.Data))
	fmt.Printf("hdr of i: %#v\n", hdri, data)

	//fmt.Println("initial slice: ", i, cap(i), len(i))
	j := append(i, 100)
	hdrj := (*reflect.SliceHeader)(unsafe.Pointer(&j))
	data1 := *(*[4]int)(unsafe.Pointer(hdrj.Data))
	fmt.Printf("hdr of j: %#v\n", hdrj, data1)

	g := append(i, 101)
	hdrg := (*reflect.SliceHeader)(unsafe.Pointer(&g))
	//data2 := *(*[4]int)(unsafe.Pointer(hdrg.Data))
	fmt.Printf("hdr of g: %#v\n", hdrg)

	h := append(i, 102)
	hdrh := (*reflect.SliceHeader)(unsafe.Pointer(&h))
	//data3 := *(*[4]int)(unsafe.Pointer(hdrh.Data))
	fmt.Printf("hdr of h: %#v\n", hdrh)

	fmt.Printf("i: %v\nj: %v\ng: %v\nh:%v\n", i, j, g, h)
}

func sliceFromLiteral() {
	fmt.Printf("\n\n** working as expected: **\n")
	// this doesnt work
	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("initial slice: ", i, cap(i), len(i))
	j := append(i, 100)
	g := append(i, 101)
	h := append(i, 102)
	fmt.Printf("i: %v\nj: %v\ng: %v\nh:%v\n", i, j, g, h)
}

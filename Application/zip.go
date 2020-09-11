//Zip functionality python in go
package main

import "fmt"

type intTuple struct {
	a, b int
}

func zip(a, b []int) ([]intTuple, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: argument must be of same length")
	}
	result := make([]intTuple, len(a), len(a))

	for i, e := range a {
		result[i] = intTuple{e, b[i]}

	}

	return result, nil
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{10, 20, 30, 40, 50, 60}

	fmt.Println(zip(a, b))
}

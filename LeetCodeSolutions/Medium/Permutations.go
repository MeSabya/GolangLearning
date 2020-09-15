package main

import "fmt"

/*
For example, {1, 2, 3} has the following six permutations:

{1, 2, 3}
{1, 3, 2}
{2, 1, 3}
{2, 3, 1}
{3, 1, 2}
{3, 2, 1}
*/

func insertAt(arr *[]int, idx int, val int) {
	//Increase the size of the slice by appending 0 or "" for string type
	*arr = append(*arr, 0)
	copy((*arr)[idx+1:], (*arr)[idx:])
	(*arr)[idx] = val
}
func permute(nums []int) [][]int {
	var result [][]int
	var permutations [][]int //intermediatory permutations
	permutations = append(permutations, []int{})

	for _, num := range nums {
		n := len(permutations)
		fmt.Println("length is ", n)
		//Now iterate over all stored permutations
		for j := 0; j < n; j++ {
			oldpermutation := permutations[0]
			permutations = permutations[1:]

			for idx := 0; idx <= len(oldpermutation); idx++ {
				newpermutation := make([]int, len(oldpermutation))
				copy(newpermutation, oldpermutation)
				insertAt(&newpermutation, idx, num)

				if len(newpermutation) == len(nums) {
					result = append(result, newpermutation)
				} else {
					permutations = append(permutations, newpermutation)
				}
			}
		}
	}

	return result
}

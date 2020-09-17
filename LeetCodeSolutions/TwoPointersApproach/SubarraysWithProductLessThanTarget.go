package main

import "fmt"

func print(nums []int) {
	for idx := 0; idx < len(nums); idx++ {
		fmt.Println(nums[idx])
	}
}

func allSubarrayProductLessThanK(nums []int, target int) {
	left := 0
	product := 1
	var result [][]int
	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		for product >= target && left < len(nums) {
			product /= nums[left]
			left++
		}

		var templist []int
		for i := right; i >= left; i-- {
			templist = append([]int{nums[i]}, templist...)
			//print(templist)
			result = append(result, templist)

		}
	}
	fmt.Println("Length of the result is ", len(result))
	for idx := 0; idx < len(result); idx++ {
		fmt.Print(result[idx])
	}
	fmt.Println()

}

func main() {
	allSubarrayProductLessThanK([]int{8, 2, 6, 5}, 50)
	allSubarrayProductLessThanK([]int{2, 5, 3, 10}, 30)

}

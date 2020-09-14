/*
subarray sum equals to K
Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.

Example 1:

Input:nums = [1,1,1], k = 2
Output: 2

Reference:
https://leetcode.com/problems/subarray-sum-equals-k/discuss/522800/Python-PrefixSum-(Logic-%2B-step-by-step-explanations)-beat-100
*/
package main

import "fmt"

func countSubarraySum(nums []int, k int) int {
	total, count := 0, 0
	prefixsum := make(map[int]int)
	prefixsum[0] = 1

	for _, num := range nums {
		total += num
		//Check val in prefixsum exists so that (total-val = k). If so increment count
		count += prefixsum[total-k]
		prefixsum[total] = prefixsum[total] + 1
	}
	return count
}

func main() {
	nums := [...]int{1, 1, 1}
	count := countSubarraySum(nums[:], 2)
	fmt.Println(count)
}

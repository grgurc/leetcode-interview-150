package main

import (
	"fmt"
	"slices"
)

func majorityElement(nums []int) int {
	slices.Sort(nums)

	for i := range (len(nums) + 1) / 2 {
		if nums[i] == nums[i+len(nums)/2] {
			return nums[i]
		}
	}

	//impossible but ok
	return -1
}

func main() {
	in := [][]int{{3, 2, 3}, {1}, {1, 1, 1, 2}}
	out := []int{3, 1, 1}

	for i := range in {
		sol := majorityElement(in[i])
		fmt.Println(sol)
		fmt.Println(out[i])
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"slices"
)

func rotate(nums []int, k int) {
	k = k % len(nums)
	if len(nums) == 0 {
		return
	}

	var firstK []int
	for i := range len(nums) - k {
		firstK = append(firstK, nums[i])
	}

	for i := len(nums) - k; i < len(nums); i++ {
		nums[i-len(nums)+k] = nums[i]
	}

	for i := k; i < len(nums); i++ {
		nums[i] = firstK[i-k]
	}
}

func main() {
	in := []int{1, 2, 3, 4, 5}
	k := 3
	out := []int{3, 4, 5, 1, 2}

	in = []int{1, 2, 3}
	k = 1
	out = []int{3, 1, 2}

	rotate(in, k)
	if slices.Compare(in, out) != 0 {
		fmt.Println("Fail")
		return
	}

	fmt.Println("Success")
}

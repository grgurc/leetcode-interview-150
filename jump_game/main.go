package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	jumpsLeft := nums[0]
	for i := 1; i < len(nums); i++ {
		jumpsLeft--
		if jumpsLeft < 0 {
			return false
		}
		jumpsLeft = max(jumpsLeft, nums[i])
	}

	return jumpsLeft >= 0
}

func main() {
	in := []int{2, 3, 1, 1, 4}
	out := true

	if canJump(in) != out {
		fmt.Println("fail")
		return
	}
	fmt.Println("success")
}

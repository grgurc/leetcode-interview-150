package main

import "fmt"

// slow solution - O(n2)
// and a lot of memory footprint - O(n)
func rec(nums []int, pos int, jumps int) int {
	if pos <= 0 {
		return jumps
	}

	minJumps := 100000
	for i := pos - 1; i >= 0; i-- {
		if nums[i]+i >= pos {
			minJumps = min(minJumps, rec(nums, i, jumps+1))
		}
	}

	return minJumps
}

func jump(nums []int) int {
	return rec(nums, len(nums)-1, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := []int{2, 3, 1, 1, 4}
	out := 2

	if jump(in) != out {
		fmt.Println(jump(in))
		return
	}
	fmt.Println("success")
}

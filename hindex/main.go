package main

import (
	"fmt"
	"slices"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func hIndex(citations []int) int {
	slices.Sort(citations)
	slices.Reverse(citations)
	maxNum := 0
	for i, k := range citations {
		maxNum = max(maxNum, min(i+1, k))
	}

	return maxNum
}

func main() {
	in := []int{3, 0, 6, 1, 5}
	out := 3
	if hIndex(in) != out {
		fmt.Println("fail")
		return
	}
	fmt.Println("success")
}

package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(prices []int) int {
	lowestPrice := 100000
	maxProfit := 0
	for _, num := range prices {
		if num < lowestPrice {
			lowestPrice = num
		} else {
			maxProfit = max(maxProfit, num-lowestPrice)
		}
	}

	return maxProfit
}

func main() {
	in := []int{7, 1, 5, 3, 6, 4}
	out := 5
	if out != maxProfit(in) {
		fmt.Println(maxProfit(in))
		fmt.Println("fail")
		return
	}

	fmt.Println("success")
}

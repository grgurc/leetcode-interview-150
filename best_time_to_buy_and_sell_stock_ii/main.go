package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	maxProfit := 0
	if len(prices) == 0 {
		return 0
	}

	lastBottom := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] { // peak, dump it
			maxProfit += max(0, prices[i-1]-lastBottom)
			lastBottom = prices[i]
		}

		if i == len(prices)-1 { // end, dump it
			maxProfit += max(0, prices[i]-lastBottom)
		}
	}

	return maxProfit
}

func main() {
	in := []int{7, 3, 4, 5, 4, 3, 1, 6}
	profit := 7

	in = []int{5, 4, 3, 2, 1}
	profit = 0
	if maxProfit(in) != profit {
		fmt.Println("fail")
		return
	}

	fmt.Println("success")
}

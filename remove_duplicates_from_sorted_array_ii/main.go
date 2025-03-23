package main

import "fmt"

func removeDuplicates(nums []int) int {
	prevNum := -100000
	prevPos := 0
	found := true
	for {
		if prevPos >= len(nums) || !found {
			break
		}
		found = false

		for i := prevPos; i < len(nums); i++ {
			if nums[i] > prevNum {
				found = true
				prevNum = nums[i]
				nums[i], nums[prevPos] = nums[prevPos], nums[i]
				prevPos++
				for j := i + 1; j < len(nums); j++ {
					if nums[j] == prevNum {
						nums[j], nums[prevPos] = nums[prevPos], nums[j]
						prevPos++
						break
					}
				}
				break
			}
		}
	}

	return prevPos
}

func main() {
	in := [][]int{{1, 1, 1, 2, 2, 3}, {}, {1, 1, 1, 2}}
	out := [][]int{{1, 1, 2, 2, 3, 1}, {}, {1, 1, 2, 1}}

	for i := range in {
		unique := removeDuplicates(in[i])
		for j := range unique {
			if in[i][j] != out[i][j] {
				fmt.Println(unique)
				fmt.Println(in[i])
				fmt.Println("Fail")
				return
			}
		}

		fmt.Println(in[i])
		fmt.Println(unique)
		fmt.Println("Success")
	}
}

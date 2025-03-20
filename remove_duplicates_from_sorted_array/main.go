package main

import "fmt"

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i+1] = nums[i+1], nums[j]
				break
			}
			if j == len(nums)-1 {
				return i + 1
			}
		}
	}
	return len(nums)
}

func main() {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	output := []int{0, 1, 2, 3, 4, 0, 1, 1, 2, 3}

	unique := removeDuplicates(input)

	for i := 0; i < unique; i++ {
		if input[i] != output[i] {
			fmt.Println("Fail")
			return
		}
	}

	fmt.Println("Success")
}

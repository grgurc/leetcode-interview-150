package removeelement

import "fmt"

func removeElement(nums []int, val int) int {
	for i := range nums {
		if nums[i] == val {
			nums[i] = -1
		}
	}

	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i] != -1 {
			i++
		}
		for j > 0 && nums[j] == -1 {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	return i
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	resSlice := []int{2, 2}

	res := removeElement(nums, val)
	if res != 2 {
		fmt.Println("fail")
		return
	}

	for i, v := range resSlice {
		if nums[i] != v {
			fmt.Println("fail")
			return
		}
	}

	fmt.Println("success")
}

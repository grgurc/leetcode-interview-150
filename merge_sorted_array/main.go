package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := n + m - 1 // start from the end of the array
	m--              // 0based index
	n--
	for pos >= 0 {
		if m < 0 {
			nums1[pos], nums2[n] = nums2[n], nums1[pos]
			n--
		} else if n < 0 {
			nums1[pos], nums1[m] = nums1[m], nums1[pos]
			m--
		} else if nums1[m] >= nums2[n] {
			nums1[pos], nums1[m] = nums1[m], nums1[pos]
			m--
		} else if nums2[n] > nums1[m] {
			nums1[pos], nums2[n] = nums2[n], nums1[pos]
			n--
		}
		pos--
	}
}

func main() {
	nums1 := []int{4, 5, 6, 7, 8, 0, 0, 0, 0}
	m := 5
	nums2 := []int{1, 2, 3, 4}
	n := 4
	merge(nums1, m, nums2, n)
	expected := []int{1, 2, 3, 4, 4, 5, 6, 7, 8}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != expected[i] {
			fmt.Println("Failed")
			fmt.Println(nums1)
			fmt.Println(nums2)
			return
		}
	}

	fmt.Println("Success")
}

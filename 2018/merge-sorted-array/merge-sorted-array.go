package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		ai, bi int
		nums3  = make([]int, 0, m+n)
	)
	for len(nums3) < m+n {
		if ai >= m {
			nums3 = append(nums3, nums2[bi])
			bi++
			continue
		}
		if bi >= n {
			nums3 = append(nums3, nums1[ai])
			ai++
			continue
		}
		if nums1[ai] < nums2[bi] {
			nums3 = append(nums3, nums1[ai])
			ai++
		} else {
			nums3 = append(nums3, nums2[bi])
			bi++
		}
	}
	for i := range nums3 {
		nums1[i] = nums3[i]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	a := []int{1, 3, 5, 0, 0}
	merge(a, 0, []int{2, 4}, 2)
	fmt.Println(a)
}

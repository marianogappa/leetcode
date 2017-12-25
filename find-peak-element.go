package main

import "fmt"

func bs(nums []int, min, max int) int {
	if min == max {
		return min
	}
	var (
		m     = (min + max) / 2
		left  = m - 1
		right = m + 1
	)
	if left < min {
		left = min
	}
	if right > max {
		right = max
	}
	switch {
	case nums[left] < nums[m] && nums[m] < nums[right]:
		return bs(nums, m, max)
	case nums[left] > nums[m] && nums[m] > nums[right]:
		return bs(nums, min, m)
	case nums[left] > nums[m] && nums[m] < nums[right]:
		return bs(nums, m, max)
	default:
		return m
	}
}

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return 0 // to avoid panic
	}
	if len(nums) == 1 {
		return 0 // because -oo < 0 < -oo
	}
	if nums[1] < nums[0] {
		return 0
	}
	if nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums) - 1
	}
	return bs(nums, 0, len(nums)-1)
}

func main() {
	var ts = []struct {
		i []int
		e int
	}{
		{i: []int{}, e: 0},
		{i: []int{1}, e: 0},
		{i: []int{2, 1}, e: 0},
		{i: []int{1, 2}, e: 1},
		{i: []int{1, 2, 3, 1}, e: 2},
		{i: []int{1, 3, 2, 1}, e: 1},
		{i: []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11}, e: 9},
	}
	for _, t := range ts {
		var a = findPeakElement(t.i)
		if t.e != a {
			fmt.Printf("findPeakElement(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

package main

import "fmt"

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
	j := len(nums) - 2
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
		if nums[j] > nums[j-1] && nums[j] > nums[j+1] {
			return j
		}
		j--
	}
	return 0 // to avoid compile error
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
		{i: []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11}, e: 9},
	}
	for _, t := range ts {
		var a = findPeakElement(t.i)
		if t.e != a {
			fmt.Printf("findPeakElement(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var (
		start, end, mid int
	)
	end = len(nums) - 1
	for start < end {
		mid = (end-start)/2 + start
		switch {
		case nums[mid] < target:
			start = mid + 1
		case nums[mid] > target:
			end = mid - 1
		default:
			start, end = mid, mid
			// This is shit! It's O(n) instead of O(log n); consider if nums is filled with target.
			// fix at some point by doing 2 binary searches.
			for i := mid; i >= 0 && nums[i] == target; i-- {
				start--
			}
			start++
			for i := mid; i < len(nums) && nums[i] == target; i++ {
				end++
			}
			end--
			return []int{start, end}
		}
	}
	if start == end && nums[start] == target {
		return []int{start, end}
	}
	return []int{-1, -1}
}

func main() {
	var ts = []struct {
		nums   []int
		target int
		e      []int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			e:      []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 9,
			e:      []int{-1, -1},
		},
		{
			nums:   []int{5},
			target: 9,
			e:      []int{-1, -1},
		},
		{
			nums:   []int{},
			target: 9,
			e:      []int{-1, -1},
		},
		{
			nums:   []int{9},
			target: 9,
			e:      []int{0, 0},
		},
		{
			nums:   []int{9, 9},
			target: 9,
			e:      []int{0, 1},
		},
		{
			nums:   []int{9, 9, 9},
			target: 9,
			e:      []int{0, 2},
		},
	}
	for _, t := range ts {
		var a = searchRange(t.nums, t.target)
		if !reflect.DeepEqual(t.e, a) {
			fmt.Printf("searchRange(%v, %v) should have been %v but was %v\n", t.nums, t.target, t.e, a)
		}
	}
}

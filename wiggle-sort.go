package main

import (
	"fmt"
)

func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		if (i%2 == 0 && nums[i-1] < nums[i]) || (i%2 == 1 && nums[i-1] > nums[i]) {
			nums[i-1], nums[i] = nums[i], nums[i-1]
		}
	}
}

func main() {
	var ts = []struct {
		i []int
	}{
		{
			i: []int{1, 2, 3},
		},
		{
			i: []int{3, 5, 2, 1, 6, 4},
		},
		{
			i: []int{3, 3},
		},
		{
			i: []int{3, 3, 3},
		},
		{
			i: []int{3, 2, 1},
		},
		{
			i: []int{3, 3, 1},
		},
		{
			i: []int{2, 1},
		},
	}
	for _, t := range ts {
		wiggleSort(t.i)
		fmt.Println("wiggleSort produced", t.i)
	}
}

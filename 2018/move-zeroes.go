package main

import (
	"fmt"
	"reflect"
)

func moveZeroes(nums []int) {
	var (
		i = 0
		j = i
	)
	for i = 0; i < len(nums); i++ {
		for j < len(nums) && nums[j] == 0 {
			j++
		}
		if j == len(nums) {
			break
		}
		nums[i] = nums[j]
		j++
	}
	for k := i; k < len(nums); k++ {
		nums[k] = 0
	}
}

func main() {
	var ts = []struct {
		i, e []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 0, 2}, []int{1, 2, 0}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{}, []int{}},
		{[]int{0, 0, 5}, []int{5, 0, 0}},
		{[]int{2, 0, 0}, []int{2, 0, 0}},
		{[]int{0, 1, 0}, []int{1, 0, 0}},
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}
	for _, t := range ts {
		moveZeroes(t.i)
		if !reflect.DeepEqual(t.e, t.i) {
			fmt.Printf("moveZeroes() should have been %v but was %v\n", t.e, t.i)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var (
		res   = make([]string, 0)
		start = 0
	)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			var newS = fmt.Sprintf("%v", nums[start])
			if i-1-start > 0 {
				newS += fmt.Sprintf("->%v", nums[i-1])
			}
			start = i
			res = append(res, newS)
		}
	}
	if start == len(nums)-1 {
		res = append(res, fmt.Sprintf("%v", nums[start]))
	} else {
		res = append(res, fmt.Sprintf("%v->%v", nums[start], nums[len(nums)-1]))
	}
	return res
}

func main() {
	var ts = []struct {
		i []int
		e []string
	}{
		{i: []int{0, 1, 2, 4, 5, 7}, e: []string{"0->2", "4->5", "7"}},
		{i: []int{0, 2, 3, 4, 6, 8, 9}, e: []string{"0", "2->4", "6", "8->9"}},
		{i: []int{}, e: []string{}},
		{i: []int{3}, e: []string{"3"}},
		{i: []int{0, 3}, e: []string{"0", "3"}},
		{i: []int{2, 3}, e: []string{"2->3"}},
		{i: []int{-1, 0}, e: []string{"-1->0"}},
		{i: []int{-2, -1}, e: []string{"-2->-1"}},
		{i: []int{-2, 0}, e: []string{"-2", "0"}},
	}
	for _, t := range ts {
		var a = summaryRanges(t.i)
		if !reflect.DeepEqual(t.e, a) {
			fmt.Printf("summaryRanges(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

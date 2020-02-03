package main

import "fmt"

func findPivot(nums []int, floor, ceil int) int {
	var n = (ceil-floor)/2 + floor
	if n > 0 && nums[n-1] > nums[n] {
		return n
	}
	if nums[n] < nums[floor] {
		return findPivot(nums, floor, n-1)
	}
	if nums[n] > nums[ceil] {
		return findPivot(nums, n+1, ceil)
	}
	if floor > 0 && nums[floor-1] > nums[floor] {
		return floor
	}
	return 0
}

func binarySearch(ns pivotedSlice, target, floor, ceil int) int {
	if ceil < floor {
		return -1
	}
	var n = (ceil-floor)/2 + floor
	switch {
	case ns.n[ns.at(n)] == target:
		return ns.at(n)
	case ns.n[ns.at(n)] < target:
		return binarySearch(ns, target, n+1, ceil)
	case ns.n[ns.at(n)] > target:
		return binarySearch(ns, target, floor, n-1)
	}
	return -1
}

type pivotedSlice struct {
	n []int
	p int
}

func (p *pivotedSlice) at(n int) int {
	if n+p.p >= len(p.n) {
		return n + p.p - len(p.n)
	}
	return n + p.p
}

func search(nums []int, target int) int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
		return -1
	}
	if len(nums) == 1 && nums[0] == target {
		return 0
	}
	return binarySearch(pivotedSlice{nums, findPivot(nums, 0, len(nums)-1)}, target, 0, len(nums)-1)
}

func main() {
	var ts = []struct {
		nums   []int
		target int
		e      int
	}{
		{nums: []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}, target: 5, e: 3},
		{nums: []int{2, 3, 1}, target: 2, e: 0},
		{nums: []int{2, 1}, target: 1, e: 1},
		{nums: []int{2, 1}, target: 5, e: -1},
		{nums: []int{10, 1}, target: 5, e: -1},
		{nums: []int{-1, -2}, target: -1, e: 0},
		{nums: []int{}, target: -1, e: -1},
		{nums: []int{1}, target: -1, e: -1},
		{nums: []int{1}, target: 1, e: 0},
		{nums: []int{1, 3}, target: 1, e: 0},
		{nums: []int{1, 3}, target: 3, e: 1},
		{nums: []int{1, 3}, target: 2, e: -1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, e: 4},
	}
	for _, t := range ts {
		var a = search(t.nums, t.target)
		if t.e != a {
			fmt.Printf("search(%v, %v) should have been %v but was %v\n", t.nums, t.target, t.e, a)
		}
	}
}

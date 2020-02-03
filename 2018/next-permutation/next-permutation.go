package main

import (
	"fmt"
	"reflect"
)

func reverse(nums []int) {
	var (
		i = 0
		j = len(nums) - 1
	)
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return
}

func qsort(a []int) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1
	pivotIndex := len(a) / 2
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
	return
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var i int
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i < 0 {
		reverse(nums)
		return
	}
	fmt.Println(i)
	for j := len(nums) - 1; j >= i+1; j-- {
		if nums[i] >= nums[j] {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		break
	}
	qsort(nums[i+1:])
}

func main() {
	var ts = []struct {
		i []int
		e []int
	}{
		{i: []int{3, 2, 1}, e: []int{1, 2, 3}},
		{i: []int{1, 1, 5}, e: []int{1, 5, 1}},
		{i: []int{1, 2, 3, 4}, e: []int{1, 2, 4, 3}},
		{i: []int{1, 2, 4, 3}, e: []int{1, 3, 2, 4}},
		{i: []int{1, 3, 2, 4}, e: []int{1, 3, 4, 2}},
		{i: []int{1, 3, 4, 2}, e: []int{1, 4, 2, 3}},
		{i: []int{1, 3, 2, 2}, e: []int{2, 1, 2, 3}},
		{i: []int{1, 3, 2, 2}, e: []int{2, 1, 2, 3}},
		{i: []int{1, 3, 2, 2}, e: []int{2, 1, 2, 3}},
		{i: []int{1, 5, 1}, e: []int{5, 1, 1}},
	}
	for _, t := range ts {
		nextPermutation(t.i)
		if !reflect.DeepEqual(t.e, t.i) {
			fmt.Println("This one didn't work", t.e, "was", t.i)
		}
	}
}

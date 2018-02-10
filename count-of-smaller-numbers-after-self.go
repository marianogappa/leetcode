package main

import (
	"fmt"
	"reflect"
)

type n struct {
	v, lc int
	l, r  *n
}

type t struct {
	r *n
}

func (t *t) insert(v int) int {
	if t.r == nil {
		t.r = &n{v, 0, nil, nil}
		return 0
	}
	return t.r.insert(v, 0)
}

func (nd *n) insert(v, c int) int {
	switch {
	case v <= nd.v:
		nd.lc++
		if nd.l == nil {
			nd.l = &n{v, 0, nil, nil}
			return c
		}
		return nd.l.insert(v, c)
	case v > nd.v:
		if nd.r == nil {
			nd.r = &n{v, 0, nil, nil}
			return c + 1 + nd.lc
		}
		return nd.r.insert(v, c+1+nd.lc)
	}
	return 0
}

func countSmaller(nums []int) []int {
	var (
		pos = make([]int, len(nums))
		t   = &t{}
	)
	for i := len(nums) - 1; i >= 0; i-- {
		pos[i] = t.insert(nums[i])
	}
	return pos
}

func main() {
	var ts = []struct {
		i []int
		e []int
	}{
		{i: []int{2, 1, 6, 2, 4, 1, 0}, e: []int{3, 1, 4, 2, 2, 1, 0}},
		{i: []int{}, e: []int{}},
		{i: []int{1}, e: []int{0}},
		{i: []int{1, 2}, e: []int{0, 0}},
		{i: []int{2, 1}, e: []int{1, 0}},
		{i: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, e: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}
	for _, t := range ts {
		var a = countSmaller(t.i)
		if !reflect.DeepEqual(t.e, a) {
			fmt.Printf("countSmaller(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

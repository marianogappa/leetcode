package main

import (
	"fmt"
	"reflect"
)

func mergesort(ns []int) []int {
	if len(ns) <= 1 {
		return ns
	}
	var mid = len(ns) / 2
	return merge(mergesort(ns[:mid]), mergesort(ns[mid:]))
}

func merge(l, r []int) []int {
	var (
		i, j int
		ns   = make([]int, len(l)+len(r))
	)
	for k := 0; k < len(l)+len(r); k++ {
		if i >= len(l) {
			ns[k] = r[j]
			j++
		} else if j >= len(r) {
			ns[k] = l[i]
			i++
		} else if l[i] < r[j] {
			ns[k] = l[i]
			i++
		} else {
			ns[k] = r[j]
			j++
		}
	}
	return ns
}

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}
	var (
		j    = 0
		res  = make([][]int, 0)
		curH = 0
		maxH int
		is   = make([]int, 0, len(buildings)*2)
		tmp  = make(map[int]struct{}, len(buildings)*2)
	)
	for i := range buildings {
		tmp[buildings[i][0]] = struct{}{}
		tmp[buildings[i][1]] = struct{}{}
	}
	for k := range tmp {
		is = append(is, k)
	}
	var sis = mergesort(is)
	for _, i := range sis {
		for j < len(buildings) && i >= buildings[j][1] {
			j++
		}
		maxH = 0
		for k := j; k < len(buildings) && buildings[k][0] <= i; k++ {
			if buildings[k][2] > maxH && i < buildings[k][1] {
				maxH = buildings[k][2]
			}
		}
		if maxH != curH {
			curH = maxH
			res = append(res, []int{i, curH})
		}
		if j >= len(buildings) {
			break
		}
	}
	return res
}

func main() {
	var ts = []struct {
		i [][]int
		e [][]int
	}{
		{
			i: [][]int{
				{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8},
			},
			e: [][]int{
				{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0},
			},
		},
		{
			i: [][]int{},
			e: [][]int{},
		},
		{
			i: [][]int{{1, 1, 1}},
			e: [][]int{},
		},
		{
			i: [][]int{{1, 2, 1}},
			e: [][]int{{1, 1}, {2, 0}},
		},
		{
			i: [][]int{{1, 2, 1}, {1, 2, 1}, {1, 2, 1}, {1, 2, 1}},
			e: [][]int{{1, 1}, {2, 0}},
		},
		{
			i: [][]int{{1, 2, 2}, {1, 2, 2}, {1, 2, 2}, {1, 2, 2}, {2, 3, 1}},
			e: [][]int{{1, 2}, {2, 1}, {3, 0}},
		},

		{
			i: [][]int{{0, 2147483647, 2147483647}},
			e: [][]int{{0, 2147483647}, {2147483647, 0}},
		},
	}
	for _, t := range ts {
		var a = getSkyline(t.i)
		if !reflect.DeepEqual(t.e, a) {
			fmt.Printf("getSkyLine(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

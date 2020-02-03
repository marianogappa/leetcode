package main

import (
	"fmt"
	"reflect"
)

func mergeSort(is []Interval) []Interval {
	if len(is) <= 1 {
		return is
	}
	var mid = len(is) / 2
	return m(mergeSort(is[mid:]), mergeSort(is[:mid]))
}

func m(l, r []Interval) []Interval {
	var (
		s      = make([]Interval, len(l)+len(r), len(l)+len(r))
		li, ri int
	)
	for k := 0; k < len(s); k++ {
		if li > len(l)-1 && ri <= len(r)-1 {
			s[k] = r[ri]
			ri++
		} else if ri > len(r)-1 && li <= len(l)-1 {
			s[k] = l[li]
			li++
		} else if l[li].Start < r[ri].Start {
			s[k] = l[li]
			li++
		} else {
			s[k] = r[ri]
			ri++
		}
	}
	return s
}

type Interval struct {
	Start int
	End   int
}

func overlap(i1, i2 Interval) (Interval, bool) {
	if i1.End >= i2.Start {
		return Interval{i1.Start, max(i1.End, i2.End)}, true
	}
	return Interval{}, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	var (
		si = mergeSort(intervals)
		fi = []Interval{si[0]}
	)
	for i := 1; i < len(si); i++ {
		if mi, ok := overlap(fi[len(fi)-1], si[i]); ok {
			fi[len(fi)-1] = mi
		} else {
			fi = append(fi, si[i])
		}
	}
	return fi
}

func main() {
	var ts = []struct {
		i []Interval
		e []Interval
	}{
		{i: []Interval{{8, 10}, {1, 3}, {2, 6}, {15, 18}}, e: []Interval{{1, 6}, {8, 10}, {15, 18}}},
		{i: []Interval{}, e: []Interval{}},
		{i: []Interval{{8, 10}}, e: []Interval{{8, 10}}},
		{i: []Interval{{1, 2}, {4, 5}}, e: []Interval{{1, 2}, {4, 5}}},
		{i: []Interval{{1, 3}, {4, 5}}, e: []Interval{{1, 3}, {4, 5}}},
		{i: []Interval{{-1, 6}, {4, 5}}, e: []Interval{{-1, 6}}},
	}
	for _, t := range ts {
		var a = merge(t.i)
		if !reflect.DeepEqual(t.e, a) {
			fmt.Printf("merge(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

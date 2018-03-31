package main

import (
	"fmt"
	"math/rand"
)

type Interval struct {
	Start int
	End   int
}

func qsort(is []int) {
	if len(is) <= 1 {
		return
	}
	var (
		l = 0
		r = len(is) - 1
		p = rand.Intn(len(is))
	)
	is[p], is[r] = is[r], is[p]
	for i := 0; i < len(is)-1; i++ {
		if is[i] < is[r] {
			is[i], is[l] = is[l], is[i]
			l++
		}
	}
	is[r], is[l] = is[l], is[r]
	qsort(is[:l])
	qsort(is[l+1:])
}

func minMeetingRooms(intervals []Interval) int {
	var (
		cur, max, i, j int
		plus, minus    = make([]int, len(intervals)), make([]int, len(intervals))
	)
	for k := range intervals {
		plus[k] = intervals[k].Start
		minus[k] = intervals[k].End
	}
	qsort(plus)
	qsort(minus)
	for i < len(plus) && j < len(minus) {
		if plus[i] < minus[j] {
			cur++
			if cur > max {
				max = cur
			}
			i++
		} else {
			cur--
			j++
		}
	}
	return max
}

func main() {
	fmt.Println(minMeetingRooms([]Interval{{0, 30}, {5, 10}, {15, 20}}) == 2)
	fmt.Println(minMeetingRooms([]Interval{{0, 1}}) == 1)
	fmt.Println(minMeetingRooms([]Interval{{0, 100}, {5, 10}, {15, 20}, {15, 20}, {15, 20}}) == 4)
}

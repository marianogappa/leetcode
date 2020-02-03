package main

import (
	"fmt"
	"math/rand"
)

type Interval struct {
	Start int
	End   int
}

// Time: O(n*logn) Space: O(n)
// If you could traverse the numbers in order (e.g. by sorting),
// every meeting time is a +1 and every meeting end is a -1 in
// "minMeetingRooms". Keep track of the global maximum.
func minMeetingRooms(intervals []Interval) int {
	var (
		cur, max, i, j int
		plus, minus    = make([]int, len(intervals)), make([]int, len(intervals))
	)

	// Time: O(n) Space: O(n)
	// Copy all ".Start"s on a "plus" slice and all ".End"s on a "minus" slice
	for k := range intervals {
		plus[k] = intervals[k].Start
		minus[k] = intervals[k].End
	}

	// Time: O(n*logn) Space: O(1)
	// Sort both slices
	quicksort(plus)
	quicksort(minus)

	// Time: O(n) Space: O(1)
	// Traverse pluses and minuses in order: each ".Start" is a +1 and each
	// ".End" is a -1
	for i < len(plus) && j < len(minus) {
		if plus[i] < minus[j] {
			cur++
			if cur > max { // Whenever you +1, update max if necessary
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

// This is vanilla quicksort; can be replaced with sorts.Ints
func quicksort(is []int) {
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
	quicksort(is[:l])
	quicksort(is[l+1:])
}

func main() {
	fmt.Println(minMeetingRooms([]Interval{{0, 30}, {5, 10}, {15, 20}}) == 2)
	fmt.Println(minMeetingRooms([]Interval{{0, 1}}) == 1)
	fmt.Println(minMeetingRooms([]Interval{{0, 100}, {5, 10}, {15, 20}, {15, 20}, {15, 20}}) == 4)
}

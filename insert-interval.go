package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

// Time: O(n*logn) Space: O(n)
// The key is to separate and sort start and end times.
// Then you can traverse in order and keep a counter:
// A start is a +1 and an end is a -1
// When the counter is 1, an interval starts.
// When the counter is 0, an interval ends.
func insert(intervals []Interval, newInterval Interval) []Interval {
	var starts, ends = make([]int, 0, len(intervals)+1), make([]int, 0, len(intervals)+1)

	// Time: O(n) Space: O(n)
	// Separate start times and end times into two sorted slices
	// Starts are already sorted except for the new interval,
	// so let's append it at the right time.
	var added bool
	for _, i := range intervals {
		if !added && newInterval.Start <= i.Start {
			starts = append(starts, newInterval.Start)
			added = true
		}
		starts = append(starts, i.Start)
		ends = append(ends, i.End)
	}
	if !added {
		starts = append(starts, newInterval.Start)
	}
	ends = append(ends, newInterval.End)

	// Time: O(n*logn) Space: O(1)
	// End times are not sorted, so we need to do it.
	sort.Ints(ends)

	// Time: O(n) Space: O(n)
	// Traverse the ordered slices of starts and ends.
	// Keeping a counter (c):
	// * A start is a +1
	// * An end is a -1
	// When the counter is 1, an interval starts.
	// When the counter is 0, an interval ends.
	var (
		s, e, c int
		newIs   = make([]Interval, 0)
	)
	for s < len(starts) {
		if starts[s] <= ends[e] {
			c++
			if c == 1 {
				newIs = append(newIs, Interval{Start: starts[s]})
			}
			s++
		} else {
			c--
			if c == 0 {
				newIs[len(newIs)-1].End = ends[e]
			}
			e++
		}
	}
	// At this point there are no further starts but there may
	// be more ends. We should take the last available end.
	newIs[len(newIs)-1].End = ends[len(ends)-1]

	return newIs
}

func main() {
	fmt.Println(insert([]Interval{{1, 3}, {6, 9}}, Interval{2, 5}))
	fmt.Println(insert([]Interval{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, Interval{4, 9}))
	fmt.Println(insert([]Interval{}, Interval{4, 9}))
	fmt.Println(insert([]Interval{{1, 5}}, Interval{2, 3}))
	fmt.Println(insert([]Interval{{1, 5}}, Interval{6, 8}))
}

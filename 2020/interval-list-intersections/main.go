package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(n)
func intervalIntersection(A [][]int, B [][]int) [][]int {
	intersections := findIntersections(A, B)
	return unionIntersections(intersections)
}

func findIntersections(A [][]int, B [][]int) [][]int {
	var (
		a, b          int
		intersections [][]int
	)
	for a < len(A) && b < len(B) {
		if isOverlap(A[a], B[b]) {
			intersections = append(intersections, intersect(A[a], B[b]))
		}
		if A[a][1] < B[b][1] {
			a++
		} else {
			b++
		}
	}
	return intersections
}

func unionIntersections(is [][]int) [][]int {
	if len(is) == 0 {
		return is
	}
	var (
		l = 0
		r = 1
	)
	for r < len(is) {
		if isOverlap(is[l], is[r]) {
			is[l] = union(is[l], is[r])
		} else {
			l++
			is[l] = is[r]
		}
		r++
	}
	return is[:l+1]
}

func isOverlap(a, b []int) bool {
	return a[0] <= b[1] && a[1] >= b[0]
}

func intersect(a, b []int) []int {
	return []int{max(a[0], b[0]), min(a[1], b[1])}
}

func union(a, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		a        [][]int
		b        [][]int
		expected [][]int
	}{
		{
			a:        [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			b:        [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			expected: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
	}
	for _, tc := range ts {
		actual := intervalIntersection(tc.a, tc.b)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.a, tc.b, tc.expected, actual)
		}
	}
}

package main

import (
	"fmt"
)

// Time: 4x3x2
// Space: O(1)
func largestTimeFromDigits(A []int) string {
	max := ""
	for i1 := 0; i1 < 4; i1++ {
		if A[i1] > 2 {
			continue
		}
		for i2 := 0; i2 < 4; i2++ {
			if i2 == i1 || (A[i1] == 2 && A[i2] > 3) {
				continue
			}
			for i3 := 0; i3 < 4; i3++ {
				if i3 == i1 || i3 == i2 || A[i3] > 5 {
					continue
				}
				for i4 := 0; i4 < 4; i4++ {
					if i4 == i1 || i4 == i2 || i4 == i3 {
						continue
					}
					Ai12 := fmt.Sprintf("%v%v", A[i1], A[i2])
					Ai34 := fmt.Sprintf("%v%v", A[i3], A[i4])
					if max != "" && (max[:2] > Ai12 || (max[:2] == Ai12 && max[3:] >= Ai34)) {
						continue
					}
					max = fmt.Sprintf("%v:%v", Ai12, Ai34)
				}
			}
		}
	}
	return max
}

func main() {
	ts := []struct {
		input    []int
		expected string
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: "23:41",
		},
		{
			input:    []int{5, 5, 5, 5},
			expected: "",
		},
		{
			input:    []int{2, 6, 3, 5},
			expected: "23:56",
		},
		{
			input:    []int{1, 1, 1, 1},
			expected: "11:11",
		},
		{
			input:    []int{0, 1, 9, 9},
			expected: "19:09",
		},
	}
	for _, tc := range ts {
		actual := largestTimeFromDigits(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

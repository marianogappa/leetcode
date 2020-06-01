package main

import "fmt"

// Time: O(1)
// Space: O(1)
func hammingWeight(num uint32) int {
	var total int
	for i := 0; i < 32; i++ {
		if num&(1<<uint(i)) > 0 {
			total++
		}
	}
	return total
}

func main() {
	ts := []struct {
		input    uint32
		expected int
	}{
		{
			input:    uint32(11),
			expected: 3,
		},
		{
			input:    uint32(128),
			expected: 1,
		},
		{
			input:    uint32(4294967293),
			expected: 31,
		},
	}
	for _, tc := range ts {
		actual := hammingWeight(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

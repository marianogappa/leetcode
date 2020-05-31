package main

import "fmt"

// Time: O(1)
// Space: O(1)
func getSum(a int, b int) int {
	for b != 0 { // while there's something to add
		c := a & b // c = carry bits
		a ^= b     // a = a+b without carrying
		b = c << 1 // b = carry bits moved 1 to the left
	}
	return a
}

func main() {
	ts := []struct {
		a        int
		b        int
		expected int
	}{
		{
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			a:        2,
			b:        3,
			expected: 5,
		},
	}
	for _, tc := range ts {
		actual := getSum(tc.a, tc.b)
		if tc.expected != actual {
			fmt.Printf("For %v + %v expected %v but got %v\n", tc.a, tc.b, tc.expected, actual)
		}
	}
}

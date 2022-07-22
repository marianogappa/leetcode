package main

import "fmt"

// Time: O(log n)
// Space: O(1)
//
// The only tricky part is calculating the complexity.
//
// The trick is thinking about how large can the sum be. For
// 3 digits, largest digits would be 999, whose square sum is
// 81*3 = 243, and for 999999999 => 81*9 = 729.
//
// So the time cannot be infinite. Either it cycles over a few hundred
// different numbers and the cycle gets found, or it drops to 1, so
// that part of the algorithm is constant.
//
// In terms of space, we only need a couple hundred items on the set,
// regardless of the size of the int, so it's constant as well.
//
// Another way to make space constant is to use slow/fast pointers to
// detect the cycle.
func isHappy(n int) bool {
	var (
		visited = map[int]struct{}{}
		sum     = n
	)
	for {
		digits := getDigits(sum)
		sum = 0
		for _, digit := range digits {
			sum += digit * digit
		}
		if sum == 1 {
			return true
		}
		if _, ok := visited[sum]; ok {
			return false
		}
		if sum <= 243 {
			visited[sum] = struct{}{}
		}
	}
}

// Time: O(log n)
// Space: O(1)
func getDigits(n int) (digits []int) {
	for n >= 10 {
		digits = append(digits, n%10)
		n /= 10
	}
	return append(digits, n)
}

func main() {
	ts := []struct {
		input    int
		expected bool
	}{
		{
			input:    19,
			expected: true,
		},
		{
			input:    2,
			expected: false,
		},
		{
			input:    1,
			expected: true,
		},
	}
	for _, tc := range ts {
		actual := isHappy(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

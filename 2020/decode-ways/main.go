package main

import "fmt"

// Time: O(n)
// Space: O(n)
func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	return doNumDecodings(s, 0, map[int]int{})
}

func doNumDecodings(s string, i int, memo map[int]int) int {
	if v, ok := memo[i]; ok {
		return v
	}
	switch len(s) - i {
	case 0:
		return 1
	case 1:
		if s[i] == '0' {
			return 0
		}
		return 1
	case 2:
		// i.e. "0[0-9]", "[3-9]0"
		if s[i] == '0' || (s[i] >= '3' && s[i+1] == '0') {
			return 0
		}
		// i.e. "10" && "20"
		if (s[i] == '1' || s[i] == '2') && s[i+1] == '0' {
			return 1
		}
		// i.e. last 2 numbers <= 26 except 10 and 20
		if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
			return 2
		}
		return 1
	default:
		// i.e. "0[0-9]", "[3-9]0"
		if s[i] == '0' || (s[i] >= '3' && s[i+1] == '0') {
			return 0
		}
		// i.e. "10" && "20"
		if (s[i] == '1' || s[i] == '2') && s[i+1] == '0' {
			memo[i] = doNumDecodings(s, i+2, memo)
			return memo[i]
		}
		// i.e. last 2 numbers <= 26 except 10 and 20
		if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
			memo[i] = doNumDecodings(s, i+1, memo) + doNumDecodings(s, i+2, memo)
			return memo[i]
		}
		memo[i] = doNumDecodings(s, i+1, memo)
		return memo[i]
	}
}

func main() {
	ts := []struct {
		input    string
		expected int
	}{
		{
			input:    "230",
			expected: 0,
		},
		{
			input:    "12",
			expected: 2,
		},
		{
			input:    "226",
			expected: 3,
		},
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "0",
			expected: 0,
		},
		{
			input:    "00",
			expected: 0,
		},
		{
			input:    "000",
			expected: 0,
		},
		{
			input:    "100",
			expected: 0,
		},
		{
			input:    "1",
			expected: 1,
		},
	}
	for _, tc := range ts {
		actual := numDecodings(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

package main

import "fmt"

// Time: O(n)
// Space: O(n)
func minRemoveToMakeValid(s string) string {
	// Do a pass left to right and a pass right to left.
	//
	// On each pass, keep a running sum while cloning
	// the string, but don't clone characters if the
	// sum would go negative.
	//
	// Reverse the string before returning (because the
	// right to left pass reversed it).
	bsLeft := []byte{}
	sum := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			sum++
			bsLeft = append(bsLeft, s[i])
		case ')':
			if sum > 0 {
				sum--
				bsLeft = append(bsLeft, s[i])
			}
		default:
			bsLeft = append(bsLeft, s[i])
		}
	}

	bsRight := []byte{}
	sum = 0
	for i := len(bsLeft) - 1; i >= 0; i-- {
		switch bsLeft[i] {
		case '(':
			if sum > 0 {
				sum--
				bsRight = append(bsRight, bsLeft[i])
			}
		case ')':
			sum++
			bsRight = append(bsRight, bsLeft[i])
		default:
			bsRight = append(bsRight, bsLeft[i])
		}
	}
	reverse(bsRight)

	return string(bsRight)
}

func reverse(bs []byte) {
	l := 0
	r := len(bs) - 1
	for l < r {
		bs[l], bs[r] = bs[r], bs[l]
		l++
		r--
	}
}

func main() {
	ts := []struct {
		input    string
		expected string
	}{
		{
			input:    "lee(t(c)o)de)",
			expected: "lee(t(c)o)de",
		},
		{
			input:    "a)b(c)d",
			expected: "ab(c)d",
		},
		{
			input:    "))((",
			expected: "",
		},
		{
			input:    "(a(b(c)d)",
			expected: "a(b(c)d)",
		},
	}
	for _, tc := range ts {
		actual := minRemoveToMakeValid(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

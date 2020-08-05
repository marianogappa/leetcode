package main

import (
	"fmt"
	"strconv"
)

// Time: O(a) where a is the number of digits in number (which is upper bounded at 9)
// Space: O(a) because we store all digits in a hashmap
func maximumSwap(number int) int {
	s := strconv.Itoa(number)

	charToLastIdx := buildCharToLastIdx(s)

	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		if bs[i] == '9' {
			continue
		}
		for _, b := range []byte{'9', '8', '7', '6', '5', '4', '3', '2', '1', '0'} {
			if b == bs[i] {
				break
			}
			idx, ok := charToLastIdx[b]
			if !ok || idx < i {
				continue
			}

			// We did one swap
			bs[i], bs[idx] = bs[idx], bs[i]
			n, _ := strconv.Atoi(string(bs))
			return n
		}
	}

	// We didn't do any swaps
	n, _ := strconv.Atoi(string(bs))
	return n
}

func buildCharToLastIdx(s string) map[byte]int {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	return m
}

func main() {
	ts := []struct {
		input    int
		expected int
	}{
		{
			input:    2736,
			expected: 7236,
		},
		{
			input:    9973,
			expected: 9973,
		},
	}
	for _, tc := range ts {
		actual := maximumSwap(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

// Time: O(n)
// Space: O(1)
func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}
	var (
		lastLetter = 0
		cursor     = 1
		count      = 1
	)
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[lastLetter] {
			count++
		} else {
			if count > 1 {
				cursor = writeCountAt(count, chars, cursor)
			}
			lastLetter = i
			count = 1
			chars[cursor] = chars[lastLetter]
			cursor++
		}
	}
	if count > 1 {
		cursor = writeCountAt(count, chars, cursor)
	}
	return cursor
}

func writeCountAt(count int, chars []byte, cursor int) int {
	number := fmt.Sprintf("%v", count)
	for j := 0; j < len(number); j++ {
		chars[cursor+j] = number[j]
	}
	return cursor + len(number)
}

func main() {
	ts := []struct {
		input         []byte
		expected      int
		expectedChars []byte
	}{
		{
			input:         []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected:      4,
			expectedChars: []byte{'a', 'b', '1', '2'},
		},
		{
			input:         []byte{},
			expected:      0,
			expectedChars: []byte{},
		},
		{
			input:         []byte{'a'},
			expected:      1,
			expectedChars: []byte{'a'},
		},
		{
			input:         []byte{'a', 'a'},
			expected:      2,
			expectedChars: []byte{'a', '2'},
		},
		{
			input:         []byte{'a', 'a', 'b'},
			expected:      3,
			expectedChars: []byte{'a', '2', 'b'},
		},
		{
			input:         []byte{'a', 'b', 'b'},
			expected:      3,
			expectedChars: []byte{'a', 'b', '2'},
		},
		{
			input:         []byte{'a', 'b', 'c'},
			expected:      3,
			expectedChars: []byte{'a', 'b', 'c'},
		},
		{
			input:         []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'a', 'b', 'c'},
			expected:      12,
			expectedChars: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', '1', '2', 'a', 'b', 'c'},
		},
	}
	for _, tc := range ts {
		actual := compress(tc.input)
		if tc.expected != actual || !reflect.DeepEqual(tc.expectedChars, tc.input[:len(tc.expectedChars)]) {
			fmt.Printf("For %s expected (%v, %s) but got %v\n", tc.input, tc.expected, tc.expectedChars, actual)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

// This exercise doesn't deserve a description.
//
// Time: O(n) because we have a single for-loop with n iterations
// Space: O(n) because we must store n strings
func fizzBuzz(n int) []string {
	var (
		m3, m5 int
		ss     = make([]string, n)
	)
	for i := 1; i <= n; i++ {
		m3++
		m5++
		switch {
		case m3 == 3 && m5 == 5:
			ss[i-1] = "FizzBuzz"
			m3 = 0
			m5 = 0
		case m3 == 3:
			ss[i-1] = "Fizz"
			m3 = 0
		case m5 == 5:
			ss[i-1] = "Buzz"
			m5 = 0
		default:
			ss[i-1] = fmt.Sprintf("%v", i)
		}
	}
	return ss
}

func main() {
	ts := []struct {
		input    int
		expected []string
	}{
		{
			input: 15,
			expected: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
		{
			input: 1,
			expected: []string{
				"1",
			},
		},
		{
			input:    0,
			expected: []string{},
		},
	}
	for _, tc := range ts {
		actual := fizzBuzz(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

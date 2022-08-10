package main

import "fmt"

// Time: O(n)
// Space: O(n)
//
// Intuitions:
// 1) It looks like some parens MUST be removed for it to be valid, e.g.:
//
// ")abc" <- this can never work
// "abc(" <- this can never work
//
// 2) Starting from the left, a running sum of "(" -> +1, ")" -> -1 detects invalids but only
//    detects them right away when they become negative.
//
// 3) By starting from the right and doing something similar but the other way around, the rest
//    of the invalids would be detected right away.
//
// So do one pass from left, one from right and don't add invalids and that's it.
func minRemoveToMakeValid(s string) string {
	fromLeft := removeInvalidsFromLeft(s)
	fromRight := removeInvalidsFromRight(fromLeft)
	return reverse(fromRight)
}

func removeInvalidsFromLeft(s string) string {
	bs := []byte{}
	runningSum := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			runningSum++
		case ')':
			if runningSum == 0 {
				continue
			}
			runningSum--
		}
		bs = append(bs, s[i])
	}
	return string(bs)
}

func removeInvalidsFromRight(s string) string {
	bs := []byte{}
	runningSum := 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case ')':
			runningSum++
		case '(':
			if runningSum == 0 {
				continue
			}
			runningSum--
		}
		bs = append(bs, s[i])
	}
	return string(bs)
}

func reverse(s string) string {
	bs := []byte(s)
	i := 0
	j := len(bs) - 1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
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
	}
	for _, tc := range ts {
		actual := minRemoveToMakeValid(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

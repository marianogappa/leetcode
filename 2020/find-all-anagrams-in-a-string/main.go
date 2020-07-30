package main

import (
	"fmt"
	"reflect"
)

// Time: O(p+s)
// Space: O(1) because the maps don't depend on the size of _s_ and _p_.
//
// This is pretty much vanilla sliding window.
// Using _windowCorrectCount_ and _windowIncorrectCount_ is
// not necessary because checking if _windowFreqs_ == _pFreqs_
// is O(1), but it's more efficient this way.
func findAnagrams(s string, p string) []int {
	var (
		start                int
		results              = []int{}
		windowFreqs          = map[byte]int{}
		windowCorrectCount   int
		windowIncorrectCount int
	)
	pFreqs := calculatePFreqs(p)
	for end := 0; end < len(s); end++ {
		windowFreqs[s[end]]++
		if f, ok := pFreqs[s[end]]; ok && windowFreqs[s[end]] <= f {
			windowCorrectCount++
		} else {
			windowIncorrectCount++
		}

		if end >= len(p) {
			if f, ok := pFreqs[s[start]]; ok && windowFreqs[s[start]] <= f {
				windowCorrectCount--
			} else {
				windowIncorrectCount--
			}
			windowFreqs[s[start]]--
			start++
		}

		if windowIncorrectCount == 0 && windowCorrectCount == len(p) {
			results = append(results, start)
		}
	}
	return results
}

func calculatePFreqs(p string) map[byte]int {
	freqs := map[byte]int{}
	for i := 0; i < len(p); i++ {
		freqs[p[i]]++
	}
	return freqs
}

func main() {
	ts := []struct {
		s        string
		p        string
		expected []int
	}{
		{
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},
		{
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
	}
	for _, tc := range ts {
		actual := findAnagrams(tc.s, tc.p)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.s, tc.p, tc.expected, actual)
		}
	}
}

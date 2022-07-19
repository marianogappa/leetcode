package main

import (
	"fmt"
	"reflect"
)

// NOTE: My solution passes but I cannot reason about the time complexity, and I don't understand the Leetcode solutions.
// Time: O(???)
// Space: O(???)

// Palindromes can come on any size >0 and there could be many partitions on the string so initially it's hard to
// think on how to traverse.
//
// All solutions must include the whole string, so one intuition would be to start by making n initial partitions at
// s[0:i+1], and if that beginning is a palindrome, then run the partition function recursively for the rest of the
// string and add as many results as this recursive invocation, always prepending the starting string to those.
//
// Two intuitions are key in this solution:
//
// 1. Even if the recursive invocation returns no results (i.e. partition("") == [][]string{}), that is still a valid
//    solution, in the case where the beginning string is the entire string, and it happens to be a palindrome.
// 2. Potentially there can be a lot of duplication of results e.g. "aaaaaaaa" where partition("aa") will be calculated
//    many times. A simple memoisation of partition results eliminate recalculation at the expense of some memory.
func partition(s string) [][]string {
	// Memoisation seems to be a good idea here, because if you consider the case "aaaaaaaa", calculating the
	// palindrome subpartitions of the first "aaaa" and last "aaaa" characters should yield the same results, so there's
	// no point calculating them twice.
	return doPartition(s, map[string][][]string{})
}

func doPartition(s string, memo map[string][][]string) [][]string {
	if s == "" {
		return [][]string{}
	}
	if res, ok := memo[s]; ok {
		return res
	}
	results := [][]string{}
	// Start by trying to partition after index i
	for i := 0; i < len(s); i++ {
		// If the string from beggining to i+1 is a palindrome
		if !isPalindrome(s[:i+1]) {
			continue
		}
		// Then calculate all subpartitions of the rest of the string, and add a partitioning result for each of those
		//
		// TRICKY POINT: note that this line MUST at least append one entry, but that should be obvious because:
		// - In the case of trying to subpartition empty string, result is empty, but that means that the entire string
		//   is a palindrome. That counts as one valid partition!
		// - In any other case, we're subpartitioning a non-empty string, so at least the 1-letter subpartitions will
		//   be palindromes.
		results = append(results, dotProduct(s[:i+1], partition(s[i+1:]))...)
	}

	memo[s] = results
	return memo[s]
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// This simply concatenates s to all subpartitions in partialResults.
// The name is technically incorrect as "dotProduct", because s is not an array.
func dotProduct(s string, partialResults [][]string) [][]string {
	// TRICKY POINT: dotProduct must return at least one result as described above, so make sure to return "s" as a
	// solution if there are no subpartitions.
	if len(partialResults) == 0 {
		return [][]string{{s}}
	}

	results := [][]string{}
	for _, partialResult := range partialResults {
		if len(partialResult) == 0 {
			continue
		}
		result := []string{s}
		result = append(result, partialResult...)

		results = append(results, result)
	}

	return results
}

func main() {
	ts := []struct {
		input    string
		expected [][]string
	}{
		{
			input:    "aab",
			expected: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			input:    "a",
			expected: [][]string{{"a"}},
		},
	}
	for _, tc := range ts {
		actual := partition(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

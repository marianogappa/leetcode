package main

import (
	"fmt"
	"strings"

	"reflect"
)

// Time: O()
// Space: O()
func wordBreak(s string, wordDict []string) []string {
	return dfs(s, wordDict, 0, map[int][]string{})
}

func dfs(s string, wordDict []string, i int, memo map[int][]string) []string {
	if i >= len(s) {
		return []string{""}
	}
	if v, ok := memo[i]; ok {
		return v
	}
	result := []string{}
	// Check which wordDicts fit at this ith index of _s_
	for _, word := range wordDict {
		if strings.HasPrefix(s[i:], word) {
			// Continue recursively after matching this word
			// and get all possible words
			recursiveSolutions := dfs(s, wordDict, i+len(word), memo)
			// Append all recursive solutions to _word_
			for _, solution := range recursiveSolutions {
				space := " "
				if solution == "" {
					space = ""
				}
				result = append(result, fmt.Sprintf("%v%v%v", word, space, solution))
			}
		}
	}

	memo[i] = result
	return result
}

func main() {
	ts := []struct {
		s        string
		wordDict []string
		expected []string
	}{
		{
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			expected: []string{"cat sand dog", "cats and dog"},
		},
		{
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			expected: []string{"pine apple pen apple",
				"pine applepen apple",
				"pineapple pen apple",
			},
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: []string{},
		},
	}
	for _, tc := range ts {
		actual := wordBreak(tc.s, tc.wordDict)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.s, tc.wordDict, tc.expected, actual)
		}
	}
}

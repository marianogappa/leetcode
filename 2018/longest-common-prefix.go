package main

import "fmt"

// Time: O(n*m) where n := len(strs) and m := len(smallest str)
// Space: O(1)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	i := doLongestCommonPrefix(strs, 0)
	if i < 0 {
		return ""
	}
	return strs[0][:i+1]
}

func doLongestCommonPrefix(strs []string, index int) int {
	first := true
	var letter byte
	for i := 0; i < len(strs); i++ {
		if index >= len(strs[i]) {
			return index - 1
		}
		if first {
			letter = strs[i][index]
			first = false
			continue
		}
		if strs[i][index] != letter {
			return index - 1
		}
	}
	return doLongestCommonPrefix(strs, index+1)
}

func main() {
	ts := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			input:    []string{"dog", "dog", ""},
			expected: "",
		},
		{
			input:    []string{"dog", "dog", "d"},
			expected: "d",
		},
		{
			input:    []string{"dog", "dog", "dog"},
			expected: "dog",
		},
		{
			input:    []string{"dog", "dog", "doggy"},
			expected: "dog",
		},
		{
			input:    []string{"dog", "cog", "fog"},
			expected: "",
		},
	}
	for _, tc := range ts {
		actual := longestCommonPrefix(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
)

// Time: O(n*k)
// Space: O(n*k) where n == len(strs) & k = len of longest string
func groupAnagrams(strs []string) [][]string {
	// Categorises strings by their letter frequencies
	categories := map[[26]int][]int{}
	for i, str := range strs {
		f := getFreq(str)
		categories[f] = append(categories[f], i)
	}

	// This horrendous thing efficiently puts
	// _categories_ into a [][]string
	res := make([][]string, len(categories))
	var i int
	for _, indices := range categories {
		cat := []string{}
		for _, index := range indices {
			cat = append(cat, strs[index])
		}
		res[i] = cat
		i++
	}

	return res
}

func getFreq(s string) [26]int {
	var f [26]int
	for i := 0; i < len(s); i++ {
		f[s[i]-'a']++
	}
	return f
}

// TODO this doesn't pass because the map makes results appear in different orders
func main() {
	ts := []struct {
		input    []string
		expected [][]string
	}{
		{
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
	}
	for _, tc := range ts {
		actual := groupAnagrams(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupStringsByAlphabet([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
}

// Time: O(n*w) where n == len(input) & w == len(largest word)
// Space: O(n*w) where n == len(input) & w == len(largest word)
func groupStringsByAlphabet(input []string) [][]string {
	groupMap := bucketInput(input)
	return mapToSliceOfSlices(groupMap)
}

func mapToSliceOfSlices(m map[string][]string) [][]string {
	results := [][]string{}
	for _, group := range m {
		result := []string{}
		for _, word := range group {
			result = append(result, word)
		}
		results = append(results, result)
	}
	return results
}

func bucketInput(input []string) map[string][]string {
	m := map[string][]string{}
	for _, word := range input {
		key := computeHash(word)
		m[key] = append(m[key], word)
	}
	return m
}

func computeHash(word string) string {
	if len(word) <= 1 {
		return ""
	}
	indices := []byte{}
	for i := 1; i < len(word); i++ {
		distance := int(word[i]) - int(word[i-1])
		if distance < 0 {
			distance += 26
		}
		indices = append(indices, byte(distance)+'a')
	}
	return string(indices)
}

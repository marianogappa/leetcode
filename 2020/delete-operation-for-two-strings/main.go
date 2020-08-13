package main

import "fmt"

// Time: O(w1*w2)
// Space: O(w1*w2)
func minDistance(word1 string, word2 string) int {
	if len(word1)*len(word2) == 0 {
		return len(word1) + len(word2)
	}

	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := range dp {
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := range word1 {
		for j := range word2 {
			if word1[i] != word2[j] {
				dp[i+1][j+1] = 1 + min(dp[i+1][j], dp[i][j+1])
			} else {
				dp[i+1][j+1] = dp[i][j]
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		word1    string
		word2    string
		expected int
	}{
		{
			word1:    "sea",
			word2:    "eat",
			expected: 2,
		},
	}
	for _, tc := range ts {
		actual := minDistance(tc.word1, tc.word2)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.word1, tc.word2, tc.expected, actual)
		}
	}
}

package main

import "fmt"

// Time: O(w1*w2)
// Space: O(w1*w2)
func minDistance(word1 string, word2 string) int {
	// If one of the words is empty, the answer is
	// inserting every character of the other one.
	if len(word1)*len(word2) == 0 {
		return len(word1) + len(word2)
	}

	// Initialising the subproblems array.
	// The content of each cell is the minimum number
	// of edits in order to go from one word to the
	// other.
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	// Going from "" to any of the words
	// should grow linearly as we're just
	// inserting.
	for i := range dp {
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	// Solve each subproblem from the base case.
	for i, b1 := range word1 {
		for j, b2 := range word2 {
			// Here's where it gets tricky:
			//
			// We're comparing the last characters of each word
			// in the subproblem (word1[:i+1], word2[:j+1]).
			//
			// If we imagine dp as a 2D matrix, to the top and left
			// of the current cell are the results before deleting or
			// inserting a character onto one of the words.
			// To the top-left is the result before replacing a
			// character.
			//
			// If the current characters are not equal, we must do 1
			// operation on top of the previous minimum result. If
			// they are equal, the result should be equal to the
			// previous minimum result.
			if b1 == b2 {
				dp[i+1][j+1] = 1 + min(dp[i][j]-1, min(dp[i][j+1], dp[i+1][j]))
			} else {
				dp[i+1][j+1] = 1 + min(dp[i][j], min(dp[i][j+1], dp[i+1][j]))
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
			word1:    "horse",
			word2:    "ros",
			expected: 3,
		},
		{
			word1:    "intention",
			word2:    "execution",
			expected: 5,
		},
		{
			word1:    "zoologicoarchaeologist",
			word2:    "zoogeologist",
			expected: 10,
		},
	}
	for _, tc := range ts {
		actual := minDistance(tc.word1, tc.word2)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.word1, tc.word2, tc.expected, actual)
		}
	}
}

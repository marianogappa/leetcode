package main

import (
	"fmt"
	"math"
)

// Time: O(n*m) where n == amount & m == len(coins)
// Space: O(n) where n == amount
func coinChange(coins []int, amount int) int {
	var minCoins = make([]int, amount+1)
	for i := 1; i < len(minCoins); i++ {
		minCoins[i] = math.MaxInt32
	}

	// I could have never figured this one out on my own.
	// Start from the bottom up. Try all coins on all
	// amounts, and as you try the larger coins, see if
	// minCoins[i-coin] has a smaller value than the
	// previous coin you tried.
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				minCoins[i] = min(minCoins[i], minCoins[i-coin]+1)
			}
		}
	}

	if minCoins[amount] == math.MaxInt32 {
		return -1
	}
	return minCoins[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		coins    []int
		amount   int
		expected int
	}{
		// {
		// 	coins:    []int{1, 2, 5},
		// 	amount:   11,
		// 	expected: 3,
		// },
		// {
		// 	coins:    []int{2},
		// 	amount:   3,
		// 	expected: -1,
		// },
		{
			coins:    []int{186, 419, 83, 408},
			amount:   6249,
			expected: 20,
		},
	}
	for _, tc := range ts {
		actual := coinChange(tc.coins, tc.amount)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.coins, tc.amount, tc.expected, actual)
		}
	}
}

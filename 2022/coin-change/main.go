package main

import (
	"fmt"
	"math"
)

// Time: O(amount*len(coins))
// Space: O(amount)
//
// Edge case #1: can we have duplicate coin denominations? No. But if so, first put the coins on a set.
// Edge case #2: can we have no coin denominations? No.
// Edge case #3: can the amount be negative? No.
// Edge case #4: can the amount be zero? YES!
//
// Reaching amount requires a "decision tree" of which coins to use. There doesn't seem to be a clear strategy for
// how to choose which coin "branch" to go for, which would require traversal of the entire tree. Whenever full
// traversal seems necessary, look for the issue of duplication in the tree. In this case, it is possible that
// different tree traversal strategies arrive at the same intermediate number, which means that there is significant
// tree duplication. When this happens, consider the possibility of memoisation, or bottom-up construction.
//
// Let's consider bottom-up:
//
// - If amount = 0, then you need 0 coins
// - If amount is equal to any of the coins, then you need one coin at the minimum.
// - If amount is larger, there are solutions if amount is the sum of a multiple of the coins, but only one solution is minimum.
//
// There doesn't seem to be a very efficient way of solving that third case.
//
// However, we can build upon the first two cases. Record those cases in a memo array, and for larger amounts, see if
// (current_amount - current_denomination) exists in the memo. If it exists, then a solution for the current amount is
// the solution for (current_amount - current_denomination) plus one coin of current_denomination. This solution may
// not be minimal, so we need to check all denominations with the current amount.
//
// Keep trying all combinations of amounts to coins until we reach the specified amount, augmenting the memo.
// By the end, we either found a minimal solution for amount, or it doesn't exist.
func coinChange(coins []int, amount int) int {
	// Optimisation for edge case #4
	if amount == 0 {
		return 0
	}

	// Initialise memo with a large number except for the base case of trying to reach 0 amount, which takes 0 coins
	minCoins := make([]int, amount+1)
	for i := range minCoins {
		minCoins[i] = math.MaxInt
	}
	minCoins[0] = 0

	// For every amount ranging 1..amount:
	for amt := 1; amt <= amount; amt++ {
		// For every coin denomination:
		for _, coin := range coins {
			// We're only interested in the cases where the coin denomination is larger than the amount (otherwise
			// cannot use it), and either this coin is exactly our amount, or the remainder is an amount for which we
			// already have calculated minCoins.
			if amt-coin < 0 || minCoins[amt-coin] == math.MaxInt {
				continue
			}

			// If we reach here, we encountered a way to reach the current amount with a number of coins, but it may
			// not be the minimum number of coins, so we must check all possible ones and keep the smallest number of
			// coins.
			minCoins[amt] = min(minCoins[amt], 1+minCoins[amt-coin])
		}
	}

	// If by the end of this exercise there isn't an option for the specified amount, then there isn't one.
	if minCoins[amount] == math.MaxInt {
		return -1
	}

	// Otherwise return it.
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
		{
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			coins:    []int{2, 5, 10, 1},
			amount:   27,
			expected: 4,
		},
	}
	for _, tc := range ts {
		actual := coinChange(tc.coins, tc.amount)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.coins, tc.amount, tc.expected, actual)
		}
	}
}

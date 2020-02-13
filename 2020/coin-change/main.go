package main

import (
	"fmt"
	"math"
)

// If we calculate how many coins it takes to add up to all amounts (from 0
// to the specified amount), the answer at each amount is the answer of an
// earlier attempt +1.
//
// At each amount goal, we look at len(coins) previous answers and +1 them,
// taking the minimum of them.
//
// We have to initialise the result of some amount goals to a large number,
// because if all coins are larger than the amount goal, we cannot use them,
// and in that case the amount goal has no solution.
//
// By the end of those _amount_ iterations, we either found the minimum
// amount of coins to reach the specified amount goal, or there isn't one.
//
// TODO: workout why the traditional DFS with memo doesn't work in this case!
//
// Time: O(n*c) where n == amount and c == len(coins)
// Space: O(n) with n == amount
func coinChange(coins []int, amount int) int {
	coinCountForAmount := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		coinCountForAmount[i] = math.MaxInt32
		for _, c := range coins {
			if c <= i {
				coinCountForAmount[i] = min(coinCountForAmount[i], coinCountForAmount[i-c]+1)
			}
		}
	}
	if coinCountForAmount[amount] == math.MaxInt32 {
		return -1
	}
	return coinCountForAmount[amount]
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
			coins:    []int{},
			amount:   0,
			expected: 0,
		},
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
			coins:    []int{186, 419, 83, 408},
			amount:   6249,
			expected: 20,
		},
	}
	for _, tc := range ts {
		actual := coinChange(tc.coins, tc.amount)
		if tc.expected != actual {
			fmt.Printf("For coins: %v, amount %v, expected %v but got %v\n", tc.coins, tc.amount, tc.expected, actual)
		}
	}
}

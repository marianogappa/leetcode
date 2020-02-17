package main

import "fmt"

// The suitable start location must have positive balance
// of gas - cost. From then on, we += further balances
// unless the += returns a negative value, in which case
// we look the next start. A running balance prior to the
// start must be kept to += at the end.
//
// Time: O(n) running through the arrays two times
// Space: O(n) to keep track of accumulated balances
func canCompleteCircuit(gas []int, cost []int) int {
	// Calculate an array of balances of gas - cost
	diffs := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		diffs[i] = gas[i] - cost[i]
	}

	// We must start from a positive balance.
	// From that start, we += all following balances.
	// If we reach a negative number, we must restart from
	// the following positive balance.
	// But when we reach the end, we'll have to += all
	// balances up to the place we started, so for keeping
	// track of balance accumulations we can reuse our
	// diffs array.
	var (
		startAccum int
		startAt    int
	)
	for i := 0; i < len(diffs); i++ {
		startAccum += diffs[i]

		// Reutilise the diffs array to keep track of
		// accumulated balances starting at diffs[0].
		if i > 0 {
			diffs[i] = diffs[i-1] + diffs[i]
		}

		// If the accumulated balances from the start
		// reach a negative value, we can't start from
		// there, so restart from the next balance.
		if startAccum < 0 {
			startAccum = 0
			startAt = i + 1
		}
	}

	// If we couldn't find a suitable start, there's no
	// solution.
	if startAt >= len(diffs) {
		return -1
	}

	// If starting from the initial value worked out,
	// there's no need to += the acummulated balances up
	// to the start index.
	if startAt == 0 {
		return startAt
	}

	// Otherwise, we must += the accumulated balance up
	// to (but not including) the start index.
	startAccum += diffs[startAt-1]

	// If +=ing that balance produces a negative result,
	// then there's no solution.
	if startAccum < 0 {
		return -1
	}

	return startAt
}

func main() {
	ts := []struct {
		gas      []int
		cost     []int
		expected int
	}{
		{
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
	}
	for _, tc := range ts {
		actual := canCompleteCircuit(tc.gas, tc.cost)
		if tc.expected != actual {
			fmt.Printf("For gas = %v, cost = %v expected %v but got %v\n", tc.gas, tc.cost, tc.expected, actual)
		}
	}
}

package main

import "fmt"

// Time: O()
// Space: O()
func canCompleteCircuit(gas []int, cost []int) int {
	diffs := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		diffs[i] = gas[i] - cost[i]
	}

	var (
		startAccum int
		startAt    int
	)
	for i := 0; i < len(diffs); i++ {
		startAccum += diffs[i]
		if i > 0 {
			diffs[i] = diffs[i-1] + diffs[i]
		}
		if startAccum < 0 {
			startAccum = 0
			startAt = i + 1
		}
	}
	if startAt >= len(diffs) {
		return -1
	}
	if startAt == 0 {
		return startAt
	}
	startAccum += diffs[startAt-1]
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

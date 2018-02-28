package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var monotonic = true
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			monotonic = false
		}
	}
	if monotonic {
		return 0
	}
	var mp = 0
	for i := 0; i < len(prices); i++ {
		for j := 0; j < len(prices); j++ {
			if i < j {
				if prices[j]-prices[i] > mp {
					mp = prices[j] - prices[i]
				}
			}
		}
	}
	return mp
}

func main() {
	ts := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
	}

	for _, t := range ts {
		if maxProfit(t.prices) != t.expected {
			fmt.Printf("maxProfit(%v) should be %v but was %v\n", t.prices, t.expected, maxProfit(t.prices))
		}
	}
}

package main

import "fmt"

func maxProfit(prices []int) int {
	var (
		maxP = 0
		buy  = 1<<31 - 1
	)
	for _, p := range prices {
		buy = min(buy, p)
		maxP = max(maxP, p-buy)
	}
	return maxP
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{1, 6, 3, 8, 4, 5, 9}))
}

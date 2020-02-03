package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	return dfs(cost, 0, make(map[int]int, 0))
}

func dfs(cs []int, i int, c map[int]int) int {
	if i+1 >= len(cs) {
		return 0
	}
	if v, ok := c[i]; ok {
		return v
	}
	c[i] = min(cs[i]+dfs(cs, i+1, c), cs[i+1]+dfs(cs, i+2, c))
	return c[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}) == 15)
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}) == 6)
	fmt.Println(minCostClimbingStairs([]int{1}) == 0)
	fmt.Println(minCostClimbingStairs([]int{1, 2}) == 1)
}

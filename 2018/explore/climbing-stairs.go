package main

import "fmt"

func climbStairs(n int) int {
	return dp(n, make(map[int]int))
}

func dp(n int, c map[int]int) int {
	if v, ok := c[n]; ok {
		return v
	}
	switch n {
	case 0, 1:
		return 1
	default:
		c[n] = dp(n-1, c) + dp(n-2, c)
		return c[n]
	}
}

func main() {
	fmt.Println(climbStairs(4000))
}

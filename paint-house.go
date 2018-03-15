package main

import "fmt"

type pos struct {
	i, lastC int
}

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	return dfs(costs, 0, -1, make(map[pos]int, 0))
}

func dfs(costs [][]int, i, lastC int, c map[pos]int) int {
	if i == len(costs) {
		return 0
	}
	if cost, ok := c[pos{i, lastC}]; ok {
		return cost
	}
	var minCost = 1<<31 - 1
	for curC := 0; curC <= 2; curC++ {
		if curC == lastC {
			continue
		}
		var cost = costs[i][curC] + dfs(costs, i+1, curC, c)
		if cost < minCost {
			minCost = cost
		}
	}
	c[pos{i, lastC}] = minCost
	return minCost
}

func main() {
	fmt.Println(minCost([][]int{
		{7, 6, 2},
	}))
}

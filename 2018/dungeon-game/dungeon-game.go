package main

import "fmt"

type xy struct {
	x, y int
}

func calculateMinimumHP(dungeon [][]int) int {
	return dfs(0, 0, dungeon, make(map[xy]int, 0))
}

func dfs(x, y int, d [][]int, c map[xy]int) int {
	if x == len(d[0])-1 && y == len(d)-1 {
		return max(-d[y][x]+1, 1)
	}
	if x >= len(d[0]) || y >= len(d) {
		return 1<<31 - 1
	}
	if r, ok := c[xy{x, y}]; ok {
		return r
	}
	var r = max(min(dfs(x+1, y, d, c), dfs(x, y+1, d, c))-d[y][x], 1)
	c[xy{x, y}] = r
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var ts = []struct {
		i [][]int
		e int
	}{
		{
			i: [][]int{
				{1},
			},
			e: 1,
		},
		{
			i: [][]int{
				{-1},
			},
			e: 2,
		},
		{
			i: [][]int{
				{-2, 3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			},
			e: 3,
		},
	}
	for _, t := range ts {
		var a = calculateMinimumHP(t.i)
		if t.e != a {
			fmt.Printf("calculateMinimumHP(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

package main

import (
	"fmt"
	"sort"
)

// This doesn't use backtracking

func combinationSum(candidates []int, target int) [][]int {
	// Time: O(n) Space: O(n)
	var ns = make([]int, 0, len(candidates))
	for _, n := range candidates {
		if n <= target {
			ns = append(ns, n)
		}
	}

	// Time: O(nlogn) Space: O(1)
	sort.Ints(ns)

	return dfs(ns, target, [][]int{{}})
}

func dfs(ns []int, target int, cur [][]int) [][]int {
	if target == 0 {
		return cur
	}
	if len(ns) == 0 {
		return [][]int{}
	}
	var newCur = make([][]int, 0)
	for i, n := range ns {
		if target-n < 0 {
			break
		}
		var curPlusN = make([][]int, 0, len(cur))
		for j := range cur {
			var tmp = make([]int, 0, len(cur[j])+1)
			tmp = append(tmp, cur[j]...)
			tmp = append(tmp, n)
			curPlusN = append(curPlusN, tmp)
		}
		newCur = append(newCur, dfs(ns[i:], target-n, curPlusN)...)
	}
	return newCur
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{1}, 1))
	fmt.Println(combinationSum([]int{1}, 10))
}

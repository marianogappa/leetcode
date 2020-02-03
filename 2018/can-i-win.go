package main

import "fmt"

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	var (
		c   = make(map[int]bool)
		sum = maxChoosableInteger * (maxChoosableInteger + 1) / 2
	)
	if desiredTotal <= 1 {
		return true
	} else if sum < desiredTotal { // if ∑numbers doesn't reach total, lose
		return false
	} else if sum == desiredTotal { // if ∑numbers == total, whoever choses last number loses
		return maxChoosableInteger%2 == 1
	}
	return dfs(maxChoosableInteger, desiredTotal, 0, c)
}

func dfs(max, tot, mask int, c map[int]bool) bool {
	if tot <= 0 { // if last play reached tot, current player loses
		return false
	}
	if v, ok := c[mask]; ok { // if we calculated before, return it
		return v
	}
	for i := 0; i < max; i++ { // try all numbers
		if mask&(1<<uint(i)) != 0 { // if number hasn't been taken
			continue
		}
		if !dfs(max, tot-i-1, mask|(1<<uint(i)), c) { // current player wins if next play loses
			c[mask] = true // store result in cache
			return true
		}
	}
	c[mask] = false // store result in cache
	return false
}

func main() {
	fmt.Println(canIWin(10, 11) == false)
	fmt.Println(canIWin(10, 10) == true)
	fmt.Println(canIWin(4, 6) == true)
	fmt.Println(canIWin(7, 16) == true)
	fmt.Println(canIWin(20, 210) == false)
}

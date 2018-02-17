package main

import "fmt"

func grayCode(n int) []int {
	var r, _ = dfs(0, map[int]struct{}{0: struct{}{}}, []int{0}, n)
	return r
}

func dfs(i int, usedM map[int]struct{}, usedS []int, n int) ([]int, bool) {
	if len(usedS) == 1<<uint(n) {
		return usedS, true
	}
	for j := 0; j < n; j++ {
		var pot = i ^ 1<<uint(j)
		if _, ok := usedM[pot]; ok {
			continue
		}
		usedM[pot] = struct{}{}
		usedS = append(usedS, pot)
		res, ok := dfs(pot, usedM, usedS, n)
		if ok {
			return res, ok
		}
		delete(usedM, pot)
		usedS = usedS[:len(usedS)]
	}
	return nil, false // to make it compile
}

func main() {
	fmt.Println(grayCode(0))
	fmt.Println(grayCode(1))
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(10))
}

package main

import "fmt"

func getFactors(n int) [][]int {
	return dp(n, 2, true)
}

func dp(n, minM int, first bool) [][]int {
	var (
		res = make([][]int, 0)
		max = n
	)
	if !first {
		max++
	}
	for i := minM; i < max; i++ {
		if n%i == 0 {
			var nn = n / i
			if nn == 1 {
				return append(res, []int{i})
			}
			var rs = dp(nn, i, false)
			for _, r := range rs {
				res = append(res, append([]int{i}, r...))
			}
		}
	}
	return res
}

func main() {
	fmt.Println(getFactors(1))
	fmt.Println(getFactors(37))
	fmt.Println(getFactors(12))
	fmt.Println(getFactors(32))
	fmt.Println(getFactors(10000))
}

package main

import "fmt"

type in struct {
	i, n int
}

func generateParenthesis(n int) []string {
	return dfs(0, n*2, make(map[in][]string, 0))
}

func dfs(i, n int, c map[in][]string) []string {
	if n == 0 {
		return []string{""}
	}
	if v, ok := c[in{i, n}]; ok {
		return v
	}
	var res = make([]string, 0)
	if i > 0 {
		var next = dfs(i-1, n-1, c)
		for _, n := range next {
			res = append(res, ")"+n)
		}
	}
	if i < n {
		var next = dfs(i+1, n-1, c)
		for _, n := range next {
			res = append(res, "("+n)
		}
	}
	c[in{i, n}] = res
	return res
}

func main() {
	fmt.Println(generateParenthesis(4))
}

package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	var (
		m = make(map[byte]struct{}, 0)
		c int
	)
	for i := 0; i < len(J); i++ {
		m[J[i]] = struct{}{}
	}
	if len(m) == 0 {
		return 0
	}
	for i := 0; i < len(S); i++ {
		if _, ok := m[S[i]]; ok {
			c++
		}
	}
	return c
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb") == 3)
	fmt.Println(numJewelsInStones("z", "ZZ") == 0)
}

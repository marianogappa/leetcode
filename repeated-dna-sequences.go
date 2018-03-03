package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	var (
		ss = make(map[int]string, 0)
		tr = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
		v  = 0
		m  = 1
	)
	// first sequence needs to be calculated completely
	for j := 9; j >= 0; j-- {
		v += tr[s[j]] * m
		m *= 4
	}
	ss[v] = ""
	// subsequent sequences only have 2 different letters
	for i := 1; i < len(s)-9; i++ {
		v = (v-tr[s[i-1]]*262144)<<2 + tr[s[i+9]]
		if _, ok := ss[v]; ok {
			ss[v] = s[i : i+10]
			continue
		}
		ss[v] = ""
	}
	// map to slice
	var (
		r = make([]string, 0)
		i = 0
	)
	for _, s := range ss {
		if len(s) > 0 {
			r = append(r, s)
			i++
		}
	}
	return r
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}

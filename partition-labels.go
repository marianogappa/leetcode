package main

import (
	"fmt"
	"reflect"
)

func partitionLabels(S string) []int { // should also make slice version; faster
	if S == "" {
		return []int{}
	}
	var ls = make(map[byte]int)
	for i := 0; i < len(S); i++ {
		ls[S[i]] = i
	}
	var (
		max     = ls[S[0]]
		from, i = 0, 0
		ps      = make([]int, 0)
	)
	for i = 0; i < len(S); i++ {
		if i <= max {
			if ls[S[i]] > max {
				max = ls[S[i]]
			}
		} else {
			ps = append(ps, i-from)
			from = i
			max = ls[S[i]]
		}
	}
	ps = append(ps, i-from)
	return ps
}

func main() {
	fmt.Println(reflect.DeepEqual(partitionLabels("ababcbacadefegdehijhklij"), []int{9, 7, 8}))
	fmt.Println(reflect.DeepEqual(partitionLabels("ababcbacadefegdehijhklija"), []int{25}))
	fmt.Println(reflect.DeepEqual(partitionLabels("a"), []int{1}))
	fmt.Println(reflect.DeepEqual(partitionLabels(""), []int{}))
}

package main

import (
	"fmt"
	"reflect"
)

func palindromePairs(words []string) [][]int {
	var ps = make([][]int, 0)
	for i := range words {
	nextCombination:
		for j := range words {
			if i != j {
				var (
					tl = len(words[i]) + len(words[j])
					ai = 0
					bi = tl - 1
				)
				for ai < bi {
					var ci, cj byte
					if ai > len(words[i])-1 {
						ci = words[j][ai-len(words[i])]
					} else {
						ci = words[i][ai]
					}
					if bi > len(words[i])-1 {
						cj = words[j][bi-len(words[i])]
					} else {
						cj = words[i][bi]
					}
					if ci != cj {
						continue nextCombination
					}
					ai++
					bi--
				}
				ps = append(ps, []int{i, j})
			}
		}
	}
	return ps
}

func main() {
	var ts = []struct {
		i []string
		e [][]int
	}{
		{
			i: []string{"bat", "tab", "cat"},
			e: [][]int{{0, 1}, {1, 0}},
		},
		{
			i: []string{"abcd", "dcba", "lls", "s", "sssll"},
			e: [][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}},
		},
		{
			i: []string{"", "aba", "ab"},
			e: [][]int{{0, 1}, {1, 0}, {2, 1}},
		},
	}
	for _, t := range ts {
		var a = palindromePairs(t.i)
		if !reflect.DeepEqual(t.e, a) {
			fmt.Printf("palindromePairs(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

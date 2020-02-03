package main

import (
	"fmt"
)

type trie struct {
	ns map[int]*trie
}

func (t *trie) insert(bs []int) {
	if len(bs) == 0 {
		return
	}
	if _, ok := t.ns[bs[0]]; !ok {
		t.ns[bs[0]] = &trie{map[int]*trie{}}
	}
	t.ns[bs[0]].insert(bs[1:])
}

func toBits(n int) []int {
	var bs = make([]int, 31)
	for i := 0; i < 31; i++ {
		bs[30-i] = (n >> uint(i)) & 1
	}
	return bs
}

func fromBits(bs []int) int {
	var n int
	for i := 0; i < 31; i++ {
		if bs[30-i] == 1 {
			n += 1 << uint(i)
		}
	}
	return n
}

func path(t1, t2 *trie) (int, int) {
	var i1, i2 int
	for i := range t1.ns {
		i1 = i
	}
	for i := range t2.ns {
		i2 = i
	}
	switch {
	case len(t1.ns)+len(t2.ns) == 4:
		return 0, 1
	case len(t1.ns)+len(t2.ns) == 2:
		return i1, i2
	case len(t1.ns) == 1:
		if i1 == 0 {
			return 0, 1
		}
		return 1, 0
	default:
		if i2 == 0 {
			return 1, 0
		}
		return 0, 1
	}
}

func resolve(t1, t2 *trie, i int, part int) int {
	if i == -1 {
		return part
	}
	var i1, i2 = path(t1, t2)
	if i1 != i2 {
		part += 1 << uint(i)
	}
	if (len(t1.ns) + len(t2.ns)) == 4 {
		return max(resolve(t1.ns[0], t2.ns[1], i-1, part), resolve(t1.ns[1], t2.ns[0], i-1, part))
	}
	return resolve(t1.ns[i1], t2.ns[i2], i-1, part)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaximumXOR(nums []int) int {
	var t = &trie{map[int]*trie{}}
	for _, n := range nums {
		t.insert(toBits(n))
	}
	return resolve(t, t, 30, 0)
}

func main() {
	var ts = []struct {
		i []int
		e int
	}{
		{i: []int{3, 10, 5, 25, 2, 8}, e: 28},
		{i: []int{1}, e: 0},
		{i: []int{8, 10, 2}, e: 10},
		{i: []int{32, 18, 33, 42, 29, 20, 26, 36, 15, 46}, e: 62},
	}
	for _, t := range ts {
		var a = findMaximumXOR(t.i)
		if t.e != a {
			fmt.Printf("findMaximumXOR(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

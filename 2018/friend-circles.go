package main

import "fmt"

type uf struct {
	v []int
	s []int
}

func newUF(n int) *uf {
	var u = uf{v: make([]int, n), s: make([]int, n)}
	for i := range u.v {
		u.v[i] = i
		u.s[i] = 1
	}
	return &u
}

func (u *uf) union(i, j int) {
	ri, rj := u.root(i), u.root(j)
	if ri == rj {
		return
	}
	if u.s[ri] < u.s[rj] {
		u.v[ri] = rj
		u.s[rj] += u.s[ri]
	} else {
		u.v[rj] = ri
		u.s[ri] += u.s[rj]
	}
}

func (u *uf) root(i int) int {
	for u.v[i] != i {
		u.v[i] = u.v[u.v[i]]
		i = u.v[i]
	}
	return i
}

func findCircleNum(M [][]int) int {
	var u = newUF(len(M))
	for i := range M {
		for j := range M[0] {
			if M[i][j] == 0 || i == j {
				continue
			}
			u.union(i, j)
		}
	}
	var rs = make(map[int]struct{}, 0)
	for i := range M {
		rs[u.root(i)] = struct{}{}
	}
	return len(rs)
}

func main() {
	var ts = []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			expected: 2,
		},
		{
			input: [][]int{
				{1, 1, 0},
				{1, 1, 1},
				{0, 1, 1},
			},
			expected: 1,
		},
		{
			input: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 3,
		},
		{
			input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 1,
		},
		{
			input: [][]int{
				{1, 0, 0, 1},
				{0, 1, 1, 0},
				{0, 1, 1, 1},
				{1, 0, 1, 1},
			},
			expected: 1,
		},
		{
			input: [][]int{
				{0},
			},
			expected: 1,
		},
		{
			input:    [][]int{},
			expected: 0,
		},
	}

	for _, t := range ts {
		if findCircleNum(t.input) != t.expected {
			fmt.Printf("findCircleNum(%v) should have been %v but was %v\n", t.input, t.expected, findCircleNum(t.input))
		}
	}
}

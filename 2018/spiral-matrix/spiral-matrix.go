package main

import (
	"fmt"
	"reflect"
)

type xy struct {
	x, y int
}

func (c1 xy) plus(c2 xy) xy {
	return xy{x: c1.x + c2.x, y: c1.y + c2.y}
}

var d = []xy{{0, 0}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func ok(c xy, my, mx int, visited map[xy]struct{}) bool {
	var _, ok = visited[c]
	return c.x >= 0 && c.y >= 0 && c.x < mx && c.y < my && !ok
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		my          = len(matrix)
		mx          = len(matrix[0])
		pos         = xy{0, 0}
		visited     = make(map[xy]struct{}, my*mx)
		res         = make([]int, 0, my*mx)
		di, retries int
	)
	for {
		if ok(pos.plus(d[di]), my, mx, visited) {
			pos = pos.plus(d[di])
			res = append(res, matrix[pos.y][pos.x])
			visited[pos] = struct{}{}
			retries = 0
		} else {
			if retries == 4 {
				return res
			}
			di++
			if di == 5 {
				di = 1
			}
			retries++
		}
	}
}

func main() {
	var ts = []struct {
		i [][]int
		e []int
	}{
		{
			i: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			e: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			i: [][]int{
				{1},
			},
			e: []int{1},
		},
		{
			i: [][]int{
				{},
			},
			e: []int{},
		},
		{
			i: [][]int{
				{1, 2},
				{4, 5},
			},
			e: []int{1, 2, 5, 4},
		},
		{
			i: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			e: []int{1, 2, 3, 6, 5, 4},
		},
		{
			i: [][]int{
				{1, 2, 3},
			},
			e: []int{1, 2, 3},
		},
		{
			i: [][]int{
				{1},
				{2},
				{3},
			},
			e: []int{1, 2, 3},
		},
	}
	for _, t := range ts {
		var a = spiralOrder(t.i)
		if !reflect.DeepEqual(t.e, a) {
			fmt.Printf("spiralOrder(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}

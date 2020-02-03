package main

import "fmt"

func bs(v func(i int) int, target, min, max int) (int, bool) {
	if min == max {
		return min, v(min) == target
	}
	var m = (max-min)/2 + min
	if min-max == 1 || min-max == -1 {
		if min > max {
			m = min
		} else {
			m = max
		}
	}
	if v(m) > target {
		if min > m-1 {
			return m, v(m) == target
		}
		return bs(v, target, min, m-1)
	} else if v(m) < target {
		if m+1 > max {
			return m, v(m) == target
		}
		return bs(v, target, m+1, max)
	} else {
		return m, true
	}
}

func getV(m [][]int, my, fy, mx, fx int) func(i int) int {
	return func(i int) int {
		return m[my*i+fy][mx*i+fx]
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var x, found = bs(getV(matrix, 0, 0, 1, 0), target, 0, len(matrix[0])-1)
	if found {
		return true
	}
	for i := x; i >= 0; i-- {
		if _, found := bs(getV(matrix, 1, 0, 0, i), target, 0, len(matrix)-1); found {
			return true
		}
	}
	return false
}

func main() {
	var ts = []struct {
		matrix [][]int
		target int
		e      bool
	}{
		{
			matrix: [][]int{
				{3, 3, 8, 13, 13, 18},
				{4, 5, 11, 13, 18, 20},
				{9, 9, 14, 15, 23, 23},
				{13, 18, 22, 22, 25, 27},
				{18, 22, 23, 28, 30, 33},
				{21, 25, 28, 30, 35, 35},
				{24, 25, 33, 36, 37, 40},
			},
			target: 21,
			e:      true,
		},
		{
			matrix: [][]int{
				{1, 6, 10, 13, 14, 16, 21},
				{3, 10, 12, 18, 22, 27, 29},
				{3, 15, 19, 20, 23, 29, 34},
				{8, 15, 19, 25, 27, 29, 39},
				{12, 17, 24, 25, 28, 29, 41},
				{16, 22, 27, 31, 31, 33, 44},
				{20, 26, 28, 35, 39, 41, 45},
				{25, 31, 34, 39, 44, 45, 47},
			},
			target: 38,
			e:      false,
		},
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			e:      true,
		},
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			e:      false,
		},
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 8,
			e:      true,
		},
		{
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 19,
			e:      true,
		},
		{
			matrix: [][]int{},
			target: 20,
			e:      false,
		},
		{
			matrix: [][]int{{}},
			target: 20,
			e:      false,
		},
		{
			matrix: [][]int{{20}},
			target: 20,
			e:      true,
		},
		{
			matrix: [][]int{{1, 1}},
			target: 2,
			e:      false,
		},
		{
			matrix: [][]int{{1}, {1}},
			target: 2,
			e:      false,
		},
		{
			matrix: [][]int{{1, 2}},
			target: 2,
			e:      true,
		},
		{
			matrix: [][]int{{2}, {1}},
			target: 2,
			e:      true,
		},
	}
	for _, t := range ts {
		var a = searchMatrix(t.matrix, t.target)
		if t.e != a {
			fmt.Printf("searchMatrix(%v, %v) should have been %v but was %v\n", t.matrix, t.target, t.e, a)
		}
	}
}

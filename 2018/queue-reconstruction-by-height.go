package main

import (
	"fmt"
	"math/rand"
)

type el struct {
	v []int
	n *el
}

type queue struct {
	h *el
}

func (q *queue) append(v []int) {
	if q.h == nil {
		q.h = &el{v, nil}
		return
	}
	var (
		e = q.h
		k = v[1]
	)
	if k == 0 {
		e = &el{v, e}
		q.h = e
		return
	}
	for e.n != nil && k > 1 {
		k--
		e = e.n
	}
	e.n = &el{v, e.n}
}

func (q *queue) dump() [][]int {
	var (
		r = make([][]int, 0)
		e = q.h
	)
	for e != nil {
		r = append(r, e.v)
		e = e.n
	}
	return r
}

func qsort(ps [][]int) {
	if len(ps) <= 1 {
		return
	}
	var (
		l = 0
		r = len(ps) - 1
		p = rand.Intn(len(ps))
	)
	ps[p], ps[r] = ps[r], ps[p]
	for i := range ps {
		if ps[i][0] > ps[r][0] || (ps[i][0] == ps[r][0] && ps[i][1] < ps[r][1]) {
			ps[l], ps[i] = ps[i], ps[l]
			if l < len(ps)-1 {
				l++
			}
		}
	}
	ps[l], ps[r] = ps[r], ps[l]
	qsort(ps[:l])
	qsort(ps[l+1:])
}

func reconstructQueue(people [][]int) [][]int {
	qsort(people)
	var q = queue{}
	for i := 0; i < len(people); i++ {
		q.append(people[i])
	}
	return q.dump()
}

func main() {
	var q = [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(q))
}

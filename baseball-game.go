package main

import "fmt"

type n struct {
	v    int
	prev *n
}

type stack struct {
	head, tail *n
}

func (s *stack) push(v int) {
	var nn = &n{v, nil}
	if s.head == nil {
		s.head, s.tail = nn, nn
		return
	}
	nn.prev, s.tail = s.tail, nn
}

func (s *stack) pop() int {
	if s.tail == nil {
		return 0
	}
	var v = s.tail.v
	s.tail = s.tail.prev
	if s.tail == nil {
		s.head = nil
	}
	return v
}

func (s *stack) sumLastTwo() int {
	if s.tail == nil {
		return 0
	}
	if s.tail.prev == nil {
		return s.tail.v
	}
	return s.tail.v + s.tail.prev.v
}

func (s *stack) doubleLast() int {
	if s.tail == nil {
		return 0
	}
	return 2 * s.tail.v
}

func readNum(s string) int {
	if len(s) == 0 {
		return 0
	}
	var (
		sign = 1
		d    = 1
		v    = 0
	)
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for i := len(s) - 1; i >= 0; i-- {
		v += int(s[i]-'0') * d
		d *= 10
	}
	return v * sign
}

func calPoints(ops []string) int {
	var (
		v, c int
		s    = stack{}
	)
	for _, o := range ops {
		switch o {
		case "+":
			v = s.sumLastTwo()
		case "D":
			v = s.doubleLast()
		case "C":
			v = s.pop()
			c -= v
			continue
		default:
			v = readNum(o)
		}
		s.push(v)
		c += v
	}
	return c
}

func main() {
	var ts = []struct {
		ops []string
		e   int
	}{
		{
			ops: []string{"5", "2", "C", "D", "+"},
			e:   30,
		},
		{
			ops: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			e:   27,
		},
		{
			ops: []string{},
			e:   0,
		},
		{
			ops: []string{"C"},
			e:   0,
		},
		{
			ops: []string{"C", "C", "C"},
			e:   0,
		},
		{
			ops: []string{"1", "2", "3", "C", "C", "C"},
			e:   0,
		},
		{
			ops: []string{"1", "2", "3", "C", "C"},
			e:   1,
		},
		{
			ops: []string{"1", "2", "+", "C", "C"},
			e:   1,
		},
		{
			ops: []string{"1", "C", "-62", "-45", "-68"},
			e:   -175,
		},
	}
	for _, t := range ts {
		var a = calPoints(t.ops)
		if t.e != a {
			fmt.Printf("calPoints(%v) should have been %v but was %v\n", t.ops, t.e, a)
		} else {
			fmt.Println("OK!")
		}
	}
}

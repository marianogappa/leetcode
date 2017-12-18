package main

import "fmt"

type li struct {
	pos        int
	prev, next *li
}

type ll struct {
	head, tail *li
}

type mi struct {
	i    *li
	dead bool
}

func (l *ll) append(pos int) *li {
	var i = &li{pos: pos}
	if l.head == nil {
		l.head = i
	}
	if l.tail == nil {
		l.tail = l.head
		return i
	}
	l.tail.next = i
	i.prev = l.tail
	l.tail = l.tail.next
	return i
}

func (l *ll) delete(i *li) {
	if l.head == i {
		l.head = l.head.next
	}
	if l.tail == i {
		if l.tail.prev != nil {
			l.tail = l.tail.prev
		} else {
			l.tail = nil
		}
	}
	if i.prev != nil {
		i.prev.next = i.next
	}
	if i.next != nil {
		i.next.prev = i.prev
	}
}

func firstUniqChar(s string) int {
	var (
		m = make(map[byte]mi, len(s))
		l ll
	)
	for i := 0; i < len(s); i++ {
		mI, ok := m[s[i]]
		if ok {
			if !mI.dead {
				mI.dead = true
				l.delete(mI.i)
				m[s[i]] = mI
			}
		} else {
			m[s[i]] = mi{i: l.append(i)}
		}
	}
	if l.head == nil {
		return -1
	}
	return l.head.pos
}

func main() {
	var ts = []struct {
		input    string
		expected int
	}{
		{input: "", expected: -1},
		{input: "leetcode", expected: 0},
		{input: "loveleetcode", expected: 2},
		{input: "aabbccdd", expected: -1},
		{input: "abcdabcd", expected: -1},
	}
	for _, t := range ts {
		var a = firstUniqChar(t.input)
		if a != t.expected {
			fmt.Printf("firstUniqChar(%v) should have been %v but was %v\n", t.input, t.expected, a)
		}
	}
}

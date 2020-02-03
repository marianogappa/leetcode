package main

import "fmt"

type node struct {
	label        int
	next, random *node
}

func copyRandomList(l *node) *node {
	if l == nil {
		return nil
	}
	var (
		m   = make(map[*node]int, 0)
		cur = l
		i   = 0
		cp  = make([]node, 0)
	)
	for cur != nil {
		m[cur] = i
		cp = append(cp, node{cur.label, nil, nil})
		i++
		cur = cur.next
	}
	cur = l
	i = 0
	for cur != nil {
		if i > 0 {
			cp[i-1].next = &cp[i]
		}
		if j, ok := m[cur.random]; ok {
			cp[i].random = &cp[j]
		}
		cur = cur.next
		i++
	}
	return &cp[0]
}

func (n *node) print() string {
	if n == nil {
		return ""
	}
	var s = fmt.Sprintf("%v", n.label)
	if n.random != nil {
		s += fmt.Sprintf("(%v)", n.random.label)
	}
	s += "->"
	return s + n.next.print()
}

func main() {
	var (
		l4 = &node{4, nil, nil}
		l3 = &node{3, l4, l4}
		l2 = &node{2, l3, l3}
		l1 = &node{1, l2, nil}
	)
	c := copyRandomList(l1)
	fmt.Println(l1.print())
	fmt.Println(c.print())
}

package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func (l *ListNode) size() int {
	if l == nil {
		return 0
	}
	return 1 + l.next.size()
}

func getIntersectionNode(l1, l2 *ListNode) *ListNode {
	var s1, s2 = l1.size(), l2.size()
	switch {
	case s1 > s2:
		var d = s1 - s2
		for i := 1; i <= d; i++ {
			l1 = l1.next
		}
	case s2 > s1:
		var d = s2 - s1
		for i := 1; i <= d; i++ {
			l2 = l2.next
		}
	}
	for l1 != nil {
		if l1 == l2 {
			return l1
		}
		l1 = l1.next
		l2 = l2.next
	}
	return nil
}

func main() {
	var sl = &ListNode{5, &ListNode{6, &ListNode{7, nil}}}
	var l1 = &ListNode{1, &ListNode{2, &ListNode{3, sl}}}
	var l2 = &ListNode{4, sl}
	fmt.Println(getIntersectionNode(l1, l2).val)

	var l3 *ListNode
	var l4 *ListNode
	fmt.Println(getIntersectionNode(l3, l4) == nil)

	var l5 = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	var l6 *ListNode
	fmt.Println(getIntersectionNode(l5, l6) == nil)

	var l7 = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	var l8 = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(getIntersectionNode(l7, l8) == nil)
}

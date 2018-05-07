package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() string {
	if l == nil {
		return ""
	}
	return fmt.Sprintf("%v->%v", l.Val, l.Next.print())
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	return doAddTwoNumbers(l1, l2, 0)
}

func doAddTwoNumbers(l1, l2 *ListNode, carry int) *ListNode {
	switch {
	case l1 == nil && l2 == nil && carry == 0:
		return nil
	case l1 == nil && l2 == nil:
		return &ListNode{carry, nil}
	case l1 == nil:
		return &ListNode{(l2.Val + carry) % 10, doAddTwoNumbers(nil, l2.Next, (l2.Val+carry)/10)}
	case l2 == nil:
		return &ListNode{(l1.Val + carry) % 10, doAddTwoNumbers(l1.Next, nil, (l1.Val+carry)/10)}
	default:
		return &ListNode{(l1.Val + l2.Val + carry) % 10, doAddTwoNumbers(l1.Next, l2.Next, (l1.Val+l2.Val+carry)/10)}
	}
}

func main() {
	var (
		l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
		l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	)
	fmt.Println(addTwoNumbers(l1, l2).print())
}

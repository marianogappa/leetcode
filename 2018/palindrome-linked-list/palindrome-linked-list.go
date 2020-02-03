package main

import "fmt"

// ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var (
		tail = head
		cur  = head
		n    = 1
	)
	if tail == nil || tail.Next == nil {
		return true
	}
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	for i := 1; i < n/2; i++ {
		cur = cur.Next
	}
	prev := cur
	cur = cur.Next
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	var (
		i = 1
		j = n
	)
	for i < j {
		if head.Val != tail.Val {
			return false
		}
		i++
		j--
		head = head.Next
		tail = tail.Next
	}
	return true
}

func listOf(is ...int) *ListNode {
	if len(is) == 0 {
		return nil
	}
	var (
		ln         = &ListNode{Val: is[0]}
		head, tail = ln, ln
	)
	for i := 1; i < len(is); i++ {
		tail.Next = &ListNode{Val: is[i]}
		tail = tail.Next
	}
	return head
}

func (l *ListNode) String() string {
	var s string
	for l != nil {
		s += fmt.Sprintf("%v->", l.Val)
		l = l.Next
	}
	return s
}

func main() {
	var ts = []struct {
		i *ListNode
		e bool
	}{
		{i: listOf(1, 2, 3), e: false},
		{i: listOf(1, 2, 1), e: true},
		{i: listOf(), e: true},
		{i: listOf(1), e: true},
		{i: listOf(1, 1), e: true},
		{i: listOf(1, 2), e: false},
		{i: listOf(1, 2, 2, 1), e: true},
		{i: listOf(1, 2, 3, 1), e: false},
	}
	for _, t := range ts {
		var a = isPalindrome(t.i)
		if t.e != a {
			fmt.Println("isPalindrome(", t.i, ") should have been", t.e, "but was", a)
		}
	}
}

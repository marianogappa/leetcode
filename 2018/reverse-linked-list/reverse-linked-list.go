package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var (
		prev = head
		tail = head.Next
		aux  *ListNode
	)
	for tail.Next != nil {
		aux = tail.Next
		tail.Next = prev
		prev = tail
		tail = aux
	}
	head.Next = nil
	tail.Next = prev
	return tail
}

func (l *ListNode) String() string {
	var s = ""
	for l != nil {
		s += fmt.Sprintf("%v->", l.Val)
		l = l.Next
	}
	return s
}

func listOf(is ...int) *ListNode {
	if len(is) == 0 {
		return nil
	}
	var (
		ln   = &ListNode{Val: is[0]}
		head = ln
		tail = ln
	)
	for i := 1; i < len(is); i++ {
		tail.Next = &ListNode{Val: is[i]}
		tail = tail.Next
	}
	return head
}

func main() {
	fmt.Println(reverseList(listOf()))
}

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		r        = &ListNode{}
		head     = r
		finished = 0
	)
	for _, l := range lists {
		if l == nil {
			finished++
		}
	}
	for finished < len(lists) {
		var (
			min  = 1<<63 - 1
			minI = 0
		)
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if lists[i].Val < min {
				min = lists[i].Val
				minI = i
			}
		}
		r.Next = &ListNode{Val: min}
		r = r.Next
		lists[minI] = lists[minI].Next
		if lists[minI] == nil {
			finished++
		}
	}

	return head.Next
}

func listOf(is ...int) *ListNode {
	var (
		r    = &ListNode{}
		head = r
	)
	for _, i := range is {
		r.Next = &ListNode{Val: i}
		r = r.Next
	}
	return head.Next
}

func main() {
	a := mergeKLists([]*ListNode{listOf(1, 3, 5), listOf(), listOf(0, 3, 10)})
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Convenient method for printing single linked lists
func (n *ListNode) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("(%v)->%v", n.Val, n.Next)
}

// In order to solve this in one pass, we must maintain a queue
// of n+1 items. This is because we need to stand on the node that
// comes right before the one to be chopped off.
//
// The simplest way I could think of maintaining this queue is with
// a "static" array of n+1 length and a rotating index to the "current"
// element in the queue. When we finish iterating the list, the target
// node (the one before the victim) will be the oldest pushed in the queue.
//
// There's an edge case: if the length of the list equals n, we cannot
// stand before the victim. In this case, just return head.Next.
//
// Time: O(n), also one pass
// Space: O(n), where n is the n from arguments + 1
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var (
		queue       = make([]*ListNode, n+1)
		cur         = head
		i, totalLen int
	)
	for cur != nil {
		totalLen++
		queue[i] = cur
		i++
		if i >= len(queue) {
			i = 0
		}
		cur = cur.Next
	}
	// special case for n == totalLen
	if totalLen == n {
		return head.Next
	}

	queue[i].Next = queue[i].Next.Next
	return head
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(removeNthFromEnd(l, 1).String() == "(1)->(2)->(3)->nil")

	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(removeNthFromEnd(l, 2).String() == "(1)->(2)->(4)->nil")

	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(removeNthFromEnd(l, 3).String() == "(1)->(3)->(4)->nil")

	l = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(removeNthFromEnd(l, 4).String() == "(2)->(3)->(4)->nil")
}

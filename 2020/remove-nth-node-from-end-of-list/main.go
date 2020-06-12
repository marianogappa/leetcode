package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%v -> %v", n.Val, n.Next)
}

// Time: O(n)
// Space: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	ahead := head
	beforeTarget := head

	// By having a pointer n steps ahead of
	// _beforeTarget_, if we move both at
	// the same time until we reach the end,
	// _beforeTarget_ will be in position to
	// remove the target node.
	for i := 0; i < n; i++ {
		ahead = ahead.Next
	}

	// Edge case: removing the head.
	if ahead == nil {
		return head.Next
	}

	// Move both pointers at the same time
	// until ahead is at the tail.
	for ahead.Next != nil {
		ahead = ahead.Next
		beforeTarget = beforeTarget.Next
	}

	// Remove target node.
	beforeTarget.Next = beforeTarget.Next.Next

	return head
}

func main() {
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 1))
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 3))
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 4))
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 5))
	fmt.Println(removeNthFromEnd(&ListNode{1, nil}, 1))
}

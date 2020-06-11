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
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	// Slow moves one node at a time, and
	// fast moves two nodes at a time.
	// If there's a cycle, they'll end up
	// together again.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	nodes := []*ListNode{
		&ListNode{3, nil},
		&ListNode{2, nil},
		&ListNode{0, nil},
		&ListNode{-4, nil},
		&ListNode{1, nil},
	}

	fmt.Println(hasCycle(&ListNode{}) == true)
}

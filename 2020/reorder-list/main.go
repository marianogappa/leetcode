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
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// 1. Find length of list
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	// 2. Divide by two (ceil)
	secondHalfIndex := length / 2
	if length%2 == 1 {
		secondHalfIndex = (length + 1) / 2
	}

	// 3. Separate second half
	beforeSecondHalf := head
	for i := 0; i < secondHalfIndex-1; i++ {
		beforeSecondHalf = beforeSecondHalf.Next
	}
	secondHalf := beforeSecondHalf.Next
	beforeSecondHalf.Next = nil // This will be the new tail, so make sure is nil

	// 4. Reverse second half
	secondHalf = reverseLinkedList(secondHalf)

	// Interleave lists
	for secondHalf != nil {
		head, secondHalf, head.Next, secondHalf.Next = head.Next, secondHalf.Next, secondHalf, head.Next
	}
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	for prev.Next != nil {
		target := prev.Next
		prev.Next = prev.Next.Next
		target.Next = head
		head = target
	}
	return head
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(l1)
	fmt.Println(l1)
	l2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(l2)
	fmt.Println(l2)
}

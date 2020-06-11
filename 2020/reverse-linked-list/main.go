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
//
// Make the drawing and think about which arrows
// need to change, including _head_. It's clearer
// to change pointers one by one, but keep in
// mind that doing it all at once is less error-
// prone.
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		target := cur.Next
		cur.Next = cur.Next.Next
		target.Next = head
		head = target
	}
	return head
}

func main() {
	fmt.Println(reverseList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}))
	fmt.Println(reverseList(&ListNode{1, nil}))
	fmt.Println(reverseList(&ListNode{1, &ListNode{2, nil}}))
	fmt.Println(reverseList(nil))
}

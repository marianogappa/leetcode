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
// Space: O(n) or O(1) if reusing l1 & l2 doesn't count as space
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		switch {
		case l1 == nil || (l2 != nil && l2.Val <= l1.Val):
			cur.Next = l2
			l2 = l2.Next
		case l2 == nil || (l1 != nil && l1.Val <= l2.Val):
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	return head.Next
}

func main() {
	fmt.Println(mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}))
	fmt.Println(mergeTwoLists(&ListNode{1, nil}, &ListNode{1, nil}))
	fmt.Println(mergeTwoLists(&ListNode{1, nil}, nil))
	fmt.Println(mergeTwoLists(nil, &ListNode{1, nil}))
	fmt.Println(mergeTwoLists(nil, nil))
}

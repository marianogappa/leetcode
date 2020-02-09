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
	return fmt.Sprintf("%v->%v", n.Val, n.Next)
}

// The strategy is to keep a pointer at the beginning
// and the last-known even and odd indexes. When we
// arrive at a new even, we link the last-known even
// to this new even (and advance the last-known), and
// the same goes to odds.
//
// By the time we reach nil, we have the lists of
// evens and odds by the head and the tail, so we
// just need to link the even tail to the odd head.
//
// Careful: the odd tail needs to be set to nil at
// the end; otherwise the final list has a cycle at
// the last node.
//
// Time: O(n)
// Space: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		curEven = head
		curOdd  = head.Next
		cur     = head.Next.Next
		headOdd = curOdd
		i       = 2
	)
	for cur != nil {
		if i%2 == 0 {
			curEven.Next = cur
			curEven = cur
		} else {
			curOdd.Next = cur
			curOdd = cur
		}
		cur = cur.Next
		i++
	}
	curEven.Next = headOdd
	curOdd.Next = nil
	return head
}

func main() {
	ts := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{
			input:    &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			expected: &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{2, &ListNode{4, nil}}}}},
		},
		{
			input:    &ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{7, nil}}}}}}},
			expected: &ListNode{2, &ListNode{3, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{5, &ListNode{4, nil}}}}}}},
		},
	}
	for _, tc := range ts {
		actual := oddEvenList(tc.input)
		if tc.expected.String() != actual.String() {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

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
	return fmt.Sprintf("(%v)->%v", n.Val, n.Next)
}

// Time: O(n log k)
// Space: O(1) or O(n) if reusing the lists counts as space
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	step := 1
	for step < len(lists) {
		// This will merge lists in pairs.
		//
		// Iteration 1: (0,1), (2,3), (4,5), ...
		// Iteration 2: (0,2), (4,6), (8,10), ...
		// Iteration 3: (0,4), (8,12), (16,20), ...
		for i := 0; i < len(lists)-step; i += step * 2 {
			// And will store the merged list into the
			// leftmost index of the "step group".
			lists[i] = merge2Lists(lists[i], lists[i+step])
		}
		step *= 2
	}
	return lists[0]
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
	head := &ListNode{} // Sentinel node
	tail := head        // Head node so we remember where we started

	// While at least one of the lists is not empty
	for list1 != nil || list2 != nil {
		switch {

		// If list2 is empty, or none is empty but list1 has the smaller Val
		case list2 == nil, list1 != nil && list2 != nil && list1.Val <= list2.Val:
			// Choose list1 and advance it
			tail.Next = list1
			list1 = list1.Next

		// If list1 is empty, or none is empty but list2 has the smaller Val
		case list1 == nil, list1 != nil && list2 != nil && list2.Val <= list1.Val:
			// Choose list2 and advance it
			tail.Next = list2
			list2 = list2.Next
		}
		// Always advance _tail_ so we can link the next node
		tail = tail.Next
	}
	tail.Next = nil  // Careful! tail.Next might be dirty!
	return head.Next // Remember the head node is the sentinel
}

func main() {
	fmt.Println(mergeKLists([]*ListNode{
		{1, &ListNode{4, &ListNode{5, nil}}},
		{1, &ListNode{3, &ListNode{4, nil}}},
		{2, &ListNode{6, nil}},
	}).String())
	fmt.Println(mergeKLists([]*ListNode{
		{1, nil},
		{1, nil},
	}).String())
	fmt.Println(mergeKLists([]*ListNode{
		nil,
		{1, nil},
	}).String())
	fmt.Println(mergeKLists([]*ListNode{
		{1, nil},
		nil,
	}).String())
	fmt.Println(mergeKLists([]*ListNode{
		nil,
		nil,
	}).String())
	// ts := []struct {
	// 	input    int
	// 	expected int
	// }{
	// 	{
	// 		input:    1,
	// 		expected: 1,
	// 	},
	// }
	// for _, tc := range ts {
	// 	actual := mergeKLists(lists []*ListNode)
	// 	if tc.expected != actual {
	// 		fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
	// 	}
	// }
}

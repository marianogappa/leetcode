package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return doAddTwoNumbers(l1, l2, 0)
}

func doAddTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil {
		if carry > 0 {
			return doAddTwoNumbers(&ListNode{carry, nil}, l2, 0)
		}
		return l2
	}
	if l2 == nil {
		if carry > 0 {
			return doAddTwoNumbers(l1, &ListNode{carry, nil}, 0)
		}
		return l1
	}
	var (
		sum   = l1.Val + l2.Val + carry
		digit = sum % 10
	)
	carry = sum / 10

	return &ListNode{digit, doAddTwoNumbers(l1.Next, l2.Next, carry)}
}

func (l *ListNode) serialize() string {
	if l == nil {
		return "nil"
	}
	return fmt.Sprintf("%v->%v", l.Val, l.Next.serialize())
}

func listOf(ns ...int) *ListNode {
	if len(ns) == 0 {
		return nil
	}
	return &ListNode{ns[0], listOf(ns[1:]...)}
}

func main() {
	fmt.Println(addTwoNumbers(listOf(2, 4, 3), listOf(5, 6, 4)).serialize())
	fmt.Println(addTwoNumbers(listOf(8, 7), listOf(9, 9, 9)).serialize())
	fmt.Println(addTwoNumbers(listOf(9, 9), listOf(0)).serialize())
	fmt.Println(addTwoNumbers(listOf(9), listOf(0)).serialize())
}

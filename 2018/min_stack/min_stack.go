package main

import "fmt"

type elem struct {
	v    int
	next *elem
	min  int
}

type MinStack struct {
	head *elem
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if this.head == nil {
		this.head = &elem{v: x, next: this.head, min: x}
		return
	}
	e := elem{v: x, next: this.head, min: min(this.head.min, x)}
	this.head = &e
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return (*this.head).v
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-1)
	fmt.Println(s.GetMin())
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.GetMin())
}

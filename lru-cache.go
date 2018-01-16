package main

import "fmt"

type el struct {
	k, v       int
	prev, next *el
}

type ll struct {
	head, tail *el
}

type LRUCache struct {
	m  map[int]*el
	l  *ll
	cp int
}

func (l *ll) evict() {
	l.tail = l.tail.prev
	l.tail.next = nil
}

func (l *ll) pop(e *el) {
	if l.head == e {
		l.head = l.head.next
	}
	if l.tail == e {
		l.tail = l.tail.prev
	}
	if e.prev != nil {
		e.prev.next = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	}
}

func (l *ll) push(e *el) {
	if l.head == nil {
		l.head, l.tail = e, e
		return
	}
	e.prev, e.next, l.head, l.head.prev = nil, l.head, e, e
}

func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]*el, 0), &ll{nil, nil}, capacity}
}

func (this *LRUCache) Get(key int) int {
	fmt.Printf("Get %v\n", key)
	var (
		e  *el
		ok bool
	)
	if e, ok = this.m[key]; !ok {
		return -1
	}
	this.l.pop(e)
	this.l.push(e)
	return e.v
}

func (this *LRUCache) Put(key int, value int) {
	fmt.Printf("Put %v: %v\n", key, value)
	if e, ok := this.m[key]; ok {
		e.v = value
		this.m[key] = e
		this.l.pop(e)
		this.l.push(e)
		return
	}
	var e = &el{key, value, nil, nil}
	this.m[key] = e
	this.l.push(e)
	if len(this.m) > this.cp {
		delete(this.m, this.l.tail.k)
		this.l.evict()
	}
}

func (e *el) print() {
	if e == nil {
		fmt.Println("nil")
		return
	}
	fmt.Printf("(%v: %v), ", e.k, e.v)
	e.next.print()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	// var aux = Constructor(2)
	// var l = &aux
	// l.Get(1)
	// l.l.head.print()
	// l.Put(1, 1)
	// l.l.head.print()
	// l.Put(2, 2)
	// l.l.head.print()
	// l.Put(3, 3)
	// l.l.head.print()
	// l.Get(3)
	// l.l.head.print()
	// l.Get(3)
	// l.l.head.print()
	// l.Get(3)
	// l.l.head.print()
	// l.Get(3)
	// l.l.head.print()
	// l.Get(3)
	// l.l.head.print()
	// l.Get(3)
	// l.l.head.print()
	// l.Get(2)
	// l.l.head.print()
	// l.Get(1)
	// l.l.head.print()
	// l.Get(0)
	// l.l.head.print()
	// l.Put(4, 4)
	// l.l.head.print()
	// l.Put(4, 5)
	// l.l.head.print()

	// var aux = Constructor(1)
	// var l = &aux
	// l.l.head.print()
	// l.Put(2, 1)
	// l.l.head.print()
	// l.Get(2)
	// l.l.head.print()
	// l.Put(3, 2)
	// l.l.head.print()
	// l.Get(2)
	// l.l.head.print()
	// l.Get(3)
	// l.l.head.print()

	// var aux = Constructor(2)
	// var l = &aux
	// l.l.head.print()
	// l.Put(2, 1)
	// l.l.head.print()
	// l.Put(1, 1)
	// l.l.head.print()
	// l.Put(2, 3)
	// l.l.head.print()
	// l.Put(4, 1)
	// l.l.head.print()
	// l.Get(1)
	// l.l.head.print()
	// l.Get(2)
	// l.l.head.print()

	// var aux = Constructor(2)
	// var l = &aux
	// l.l.head.print()
	// l.Put(2, 1)
	// l.l.head.print()
	// l.Put(2, 2)
	// l.l.head.print()
	// l.Get(2)
	// l.l.head.print()
	// l.Put(1, 1)
	// l.l.head.print()
	// l.Put(4, 1)
	// l.l.head.print()
	// l.Get(2)
	// l.l.head.print()

	var aux = Constructor(3)
	var l = &aux
	l.l.head.print()
	l.Put(1, 1)
	l.l.head.print()
	l.Put(2, 2)
	l.l.head.print()
	l.Put(3, 3)
	l.l.head.print()
	l.Put(4, 4)
	l.l.head.print()
	fmt.Println(l.Get(4))
	l.l.head.print()
	fmt.Println(l.Get(3))
	l.l.head.print()
	fmt.Println(l.Get(2))
	l.l.head.print()
	fmt.Println(l.Get(1))
	l.l.head.print()
	l.Put(5, 5)
	l.l.head.print()
	fmt.Println(l.Get(1))
	l.l.head.print()
	fmt.Println(l.Get(2))
	l.l.head.print()
	fmt.Println(l.Get(3))
	l.l.head.print()
	fmt.Println(l.Get(4))
	l.l.head.print()
	fmt.Println(l.Get(5))
	l.l.head.print()
}

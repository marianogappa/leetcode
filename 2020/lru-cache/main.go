package main

import "fmt"

type node struct {
	key, val   int
	next, prev *node
}

func (n *node) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%v->%v", n.val, n.next)
}

type dll struct {
	head, tail *node
}

func (l *dll) promote(nd *node) {
	if l.head == nd {
		return
	}
	if l.tail == nd {
		l.pushFront(l.popBack())
		return
	}
	nd.prev.next = nd.next
	nd.next.prev = nd.prev
	l.pushFront(nd)
}

func (l *dll) pushFront(nd *node) {
	nd.next = l.head
	nd.prev = nil
	if l.head != nil {
		l.head.prev = nd
	}
	l.head = nd
	if l.tail == nil {
		l.tail = nd
	}
}

func (l *dll) popBack() *node {
	target := l.tail

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return target
	}
	if target.prev != nil {
		target.prev.next = nil
		l.tail = target.prev
	} else {
		l.tail = nil
	}
	return target
}

type LRUCache struct {
	capacity int
	hash     map[int]*node
	list     *dll
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*node{}, &dll{nil, nil}}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.hash[key]; !ok {
		return -1
	}
	this.list.promote(this.hash[key])
	return this.hash[key].val
}

func (this *LRUCache) Put(key int, value int) {
	if nd, ok := this.hash[key]; ok {
		nd.val = value
		this.list.promote(nd)
	} else {
		nd := &node{key, value, nil, nil}
		this.list.pushFront(nd)
		this.hash[key] = nd
		if len(this.hash) > this.capacity {
			tl := this.list.popBack()
			delete(this.hash, tl.key)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	c := Constructor(2)
	cache := &c

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
	fmt.Println(cache.list.head)
}

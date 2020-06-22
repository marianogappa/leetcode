package main

import "fmt"

type node struct {
	next [26]*node
	stop bool
}

func (n *node) insert(s string) {
	if s == "" {
		n.stop = true
		return
	}
	c := s[0] - 'a'
	if n.next[c] == nil {
		n.next[c] = &node{next: [26]*node{}}
	}
	n.next[c].insert(s[1:])
}

func (n *node) search(s string, isStop bool) bool {
	if s == "" {
		return !isStop || n.stop
	}
	if s[0] != '.' {
		c := s[0] - 'a'
		if n.next[c] == nil {
			return false
		}
		return n.next[c].search(s[1:], isStop)
	}
	for _, n := range n.next {
		if n != nil && n.search(s[1:], isStop) {
			return true
		}
	}
	return false
}

type WordDictionary struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{root: &node{
		next: [26]*node{},
		stop: false,
	}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.root.insert(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.root.search(word, true)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

// Time: O(n^2) with all . and matching the very last branch of dfs
// Space: O(m) m == total of letters assumming no reusing of branches
func main() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	fmt.Println(obj.Search("pad") == false)
	fmt.Println(obj.Search("bad") == true)
	fmt.Println(obj.Search(".ad") == true)
	fmt.Println(obj.Search("b..") == true)
}

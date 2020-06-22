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
	c := s[0] - 'a'
	if n.next[c] == nil {
		return false
	}
	return n.next[c].search(s[1:], isStop)
}

// Time: O()
// Space: O()
type Trie struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: &node{next: [26]*node{}, stop: false}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.root.insert(word)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.root.search(word, true)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.root.search(prefix, false)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple") == true)   // returns true
	fmt.Println(trie.Search("app") == false)    // returns false
	fmt.Println(trie.StartsWith("app") == true) // returns true
	trie.Insert("app")
	fmt.Println(trie.Search("app") == true) // returns true

}

package main

import "fmt"

type Trie struct {
	isEnd bool
	m     []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{false, make([]*Trie, 26)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		this.isEnd = true
		return
	}
	var i = int(word[0] - 'a')
	if this.m[i] == nil {
		var c = Constructor()
		this.m[i] = &c
	}
	this.m[i].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.Traverse(word, false)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.Traverse(prefix, true)
}

func (this *Trie) Traverse(s string, isPrefix bool) bool {
	if s == "" {
		return isPrefix || this.isEnd
	}
	var i = int(s[0] - 'a')
	if this.m[i] == nil {
		return false
	}
	return this.m[i].Traverse(s[1:], isPrefix)
}

/**
Your Trie object will be instantiated and called as such:
obj := Constructor();
obj.Insert(word);
param_2 := obj.Search(word);
param_3 := obj.StartsWith(prefix);
*/

func main() {
	c := Constructor()
	obj := &c
	obj.Insert("hello")
	fmt.Println(obj.Search("hell"))
	fmt.Println(obj.StartsWith("hell"))
}

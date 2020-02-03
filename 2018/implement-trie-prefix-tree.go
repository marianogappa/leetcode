package main

import (
	"fmt"
)

type Trie struct {
	m           map[byte]*Trie
	isEndOfWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{map[byte]*Trie{}, false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		this.isEndOfWord = true
		return
	}
	if _, ok := this.m[word[0]]; !ok {
		this.m[word[0]] = &Trie{map[byte]*Trie{}, len(word) == 1}
	}
	this.m[word[0]].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.doSearch(word, false)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.doSearch(prefix, true)
}

func (this *Trie) doSearch(s string, partial bool) bool {
	if s == "" {
		return this.isEndOfWord || partial
	}
	t, ok := this.m[s[0]]
	if !ok {
		return false
	}
	return t.doSearch(s[1:], partial)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	obj := Constructor()
	o := &obj
	o.Insert("abc")
	if !o.Search("abc") {
		fmt.Println("FAIL: should have found abc")
	}
	if o.Search("ab") {
		fmt.Println("FAIL: should not have found ab")
	}
	o.Insert("ab")
	if !o.Search("ab") {
		fmt.Println("FAIL: should have found ab")
	}
	o.Insert("ab")
	if !o.Search("ab") {
		fmt.Println("FAIL: should have found ab")
	}
	if !o.StartsWith("a") {
		fmt.Println("FAIL: should have found a prefix")
	}
	if !o.StartsWith("ab") {
		fmt.Println("FAIL: should have found ab prefix")
	}
	if !o.StartsWith("abc") {
		fmt.Println("FAIL: should have found abc prefix")
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	m map[int]int
	i []int
	s *rand.Rand
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int, 0), make([]int, 0), rand.New(rand.NewSource(time.Now().Unix()))}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.m[val] = len(this.m)
		this.i = append(this.i, val)
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.m[val]; ok {
		delete(this.m, val)
		if i < len(this.i)-1 {
			var lastVal = this.i[len(this.i)-1]
			this.i[i] = lastVal
			this.m[lastVal] = i
		}
		this.i = this.i[:len(this.i)-1]
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.i) > 0 {
		return this.i[this.s.Intn(len(this.i))]
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	var o = Constructor()
	o.Insert(1)
	o.Insert(2)
	o.Insert(3)
	o.Insert(4)
	o.Remove(2)
	o.Remove(3)
	o.Remove(4)
	o.Remove(1)
	o.Insert(1)
	o.Insert(2)
	o.Insert(3)
	o.Insert(4)
	for i := 0; i < 500000; i++ {
		fmt.Println(o.GetRandom())
	}
}

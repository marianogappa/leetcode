package main

import "fmt"

type TwoSum struct {
	m map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{make(map[int]int, 0)}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.m[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for n := range this.m {
		this.m[n]--
		if v, ok := this.m[value-n]; ok && v > 0 {
			return true
		}
		this.m[n]++
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */

func main() {
	var o = Constructor()
	o.Add(0)
	o.Add(0)
	fmt.Println(o.Find(0))
}

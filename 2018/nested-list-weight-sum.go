package main

type NestedInteger struct {
}

func (n NestedInteger) GetInteger() int           { return 0 }
func (n NestedInteger) IsInteger() bool           { return true }
func (n NestedInteger) GetList() []*NestedInteger { return nil }

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func depthSum(nestedList []*NestedInteger) int {
	var s = 0
	for _, i := range nestedList {
		s += sum(*i, 1)
	}
	return s
}

func sum(i NestedInteger, multiplier int) int {
	if i.IsInteger() {
		return i.GetInteger() * multiplier
	}
	var s = 0
	for _, j := range i.GetList() {
		s += sum(*j, multiplier+1)
	}
	return s
}

func main() {

}

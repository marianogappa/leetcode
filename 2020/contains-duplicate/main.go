package main

import "fmt"

// Uh...hashmap?
//
// Time: O(n)
// Space: O(n)
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3}) == false)
	fmt.Println(containsDuplicate([]int{1, 2, 1}) == true)
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) == true)
}

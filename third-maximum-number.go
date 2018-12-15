package main

import "fmt"

// Time: O(n)
// Space: O(1)
func thirdMax(nums []int) int {
	var m1, m2, m3, count int
	for _, num := range nums {
		if (count >= 1 && num == m1) || (count >= 2 && num == m2) || (count >= 3 && num == m3) {
			continue
		}
		if count == 0 || num > m1 {
			m1, m2, m3 = num, m1, m2
		} else if count == 1 || num > m2 {
			m2, m3 = num, m2
		} else if count == 2 || num > m3 {
			m3 = num
		}
		count++
	}
	if count >= 3 {
		return m3
	}
	return m1
}

func main() {
	ts := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{}, 0},
		{[]int{2, 2, 2}, 2},
		{[]int{1, 2, -2147483648}, -2147483648},
		{[]int{2, 2, 2, 1}, 2},
		{[]int{3, 3, 4, 3, 4, 3, 0, 3, 3}, 0},
	}
	for _, tc := range ts {
		actual := thirdMax(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}

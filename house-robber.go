package main

import "fmt"

func rob(nums []int) int {
	var c = make([][]int, len(nums))
	for i := range c {
		c[i] = make([]int, len(nums))
		for j := range c[i] {
			c[i][j] = -1
		}
	}
	return rc(0, len(nums)-1, c, nums)
}

func rc(l, r int, c [][]int, ns []int) int {
	if l > r {
		return 0
	}
	if v := c[l][r]; v > -1 {
		return v
	}
	var max = 0
	for i := l; i <= r; i++ {
		var tot = ns[i] + rc(l, i-2, c, ns) + rc(i+2, r, c, ns)
		if tot > max {
			max = tot
		}
	}
	c[l][r] = max
	return max
}

func main() {
	var ts = []struct {
		nums []int
		e    int
	}{
		{
			nums: []int{1, 2, 3},
			e:    4,
		},
		{
			nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			e:    0,
		},
	}
	for _, t := range ts {
		var a = rob(t.nums)
		if t.e != a {
			fmt.Printf("rob(%v) should have been %v but was %v\n", t.nums, t.e, a)
		}
	}
}

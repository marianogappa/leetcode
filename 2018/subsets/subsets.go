package main

import "fmt"

func subsets(nums []int) [][]int {
	var ss = [][]int{{}}
	for _, n := range nums {
		for i := range ss {
			var s = make([]int, 0)
			s = append(s, ss[i]...)
			s = append(s, n)
			ss = append(ss, s)
		}
	}
	return ss
}

func main() {
	var ts = [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
	}
	for _, t := range ts {
		fmt.Println("for", t, "it's", subsets(t))
	}

}

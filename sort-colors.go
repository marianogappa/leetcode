package main

import "fmt"

func sortColors(nums []int) {
	var n0, n1, n2 int
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 2:
			n2++
		case 1:
			if n2 != 0 {
				nums[n0+n1], nums[i] = nums[i], nums[n0+n1]
				n2--
				i--
			}
			n1++
		case 0:
			if !(n1 == 0 && n2 == 0) {
				nums[n0], nums[i] = nums[i], nums[n0]
				if n1 > 0 {
					n1--
				} else {
					n2--
				}
				i--
			}
			n0++
		}
	}
}

func main() {
	var ts = [][]int{
		[]int{0, 1, 2, 0, 1, 2},
		[]int{2, 2, 2, 2, 2, 2},
		[]int{},
		[]int{1},
		[]int{2},
		[]int{0},
		[]int{0, 1},
		[]int{1, 0, 0},
		[]int{1, 2, 0},
		[]int{2, 1, 1},
		[]int{2, 2, 1},
	}
	for i := range ts {
		fmt.Println("from", ts[i])
		sortColors(ts[i])
		fmt.Println("to", ts[i])
	}
}

package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	var (
		c     = 0
		first = true
	)
	for _, f := range flowerbed {
		if f == 0 {
			c++
		} else {
			if first {
				if c >= 2 {
					n -= c / 2
				}
			} else {
				if c >= 3 {
					n -= (c - 1) / 2
				}
			}
			c = 0
			first = false
		}
		if n == 0 {
			return true
		}
	}
	if first {
		if c > 0 {
			n -= (c + 1) / 2
		}
	} else {
		if c >= 2 {
			n -= c / 2
		}
	}
	return n <= 0
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2) == false)
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0, 1}, 2) == true)
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 1}, 2) == false)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 1}, 1) == false)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 1, 0, 0, 1}, 1) == false)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0, 0, 1}, 2) == true)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0, 0, 1}, 1) == true)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0}, 3) == true)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0, 0, 1, 0}, 3) == false)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0}, 3) == false)
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0}, 3) == true)
}

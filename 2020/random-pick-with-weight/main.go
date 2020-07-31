package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1. Store the incremental sum in _accW_.
// 2. Pick a random number in [0,accW[len(accW)-1]].
// 3. Do binary search to find its index.
type Solution struct {
	accW []int
}

// Time: O(n)
// Space: O(n)
func Constructor(w []int) Solution {
	accW := make([]int, len(w))
	accW[0] = w[0]
	for i := 1; i < len(w); i++ {
		accW[i] = accW[i-1] + w[i]
	}
	return Solution{accW}
}

// Time: O(logn)
// Space: O(1)
func (this *Solution) PickIndex() int {
	randomVal := rand.Intn(this.accW[len(this.accW)-1]) + 1
	return this.binarySearch(randomVal, 0, len(this.accW)-1)
}

// Time: O(logn)
// Space: O(1)
func (this *Solution) binarySearch(n, left, right int) int {
	if left == right {
		return left
	}
	mid := (left + right) / 2
	if this.accW[mid] < n {
		return this.binarySearch(n, mid+1, right)
	}
	return this.binarySearch(n, left, mid)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

func main() {
	rand.Seed(time.Now().UnixNano())
	obj := Constructor([]int{2, 8})
	countZeroes := 0
	countOnes := 0
	for i := 0; i < 1000000; i++ {
		if obj.PickIndex() == 0 {
			countZeroes++
		} else {
			countOnes++
		}
	}
	fmt.Println(countZeroes, countOnes)
}

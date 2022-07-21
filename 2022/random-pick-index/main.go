package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Time: O(n)
// Space: O(n)
//
// As long as we know how many duplicates there are for each number, we can use a built-in random number generator
// (which is a constant time, constant space syscall) for that length. So simply put all indices of every number into
// a hashmap (linear time, linear space), where the key is the number, and the value is an array of indices.
type Solution struct {
	targetToIdxs map[int][]int
}

func Constructor(nums []int) Solution {
	targetToIdxs := map[int][]int{}
	for i, num := range nums {
		targetToIdxs[num] = append(targetToIdxs[num], i)
	}

	return Solution{targetToIdxs: targetToIdxs}
}

func (this *Solution) Pick(target int) int {
	availableIdxs := this.targetToIdxs[target]
	return availableIdxs[rand.Intn(len(availableIdxs))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

func main() {
	obj := Constructor([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})

	targetToProbability := map[int]float64{
		1: 1,
		2: 0.5,
		3: 0.33,
		4: 0.25,
	}
	attempts := 1000000
	for target, expectedProbability := range targetToProbability {
		results := map[int]int{}
		for i := 0; i < attempts; i++ {
			idx := obj.Pick(target)
			results[idx]++
		}
		for idx, timesPicked := range results {
			actualProbability := float64(timesPicked) / float64(attempts)
			differenceToExpectedProbability := math.Abs(actualProbability - expectedProbability)
			if differenceToExpectedProbability > 0.01 {
				fmt.Printf("For target %v, the idx %v was picked with %v probability, but %v was expected\n", target, idx, actualProbability, expectedProbability)
			}
		}
	}
}

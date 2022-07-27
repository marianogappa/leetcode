package main

import (
	"fmt"
	"math"
	"math/rand"
)

// The only tricky part on this exercise is figuring out that it would take too much space to make a mapping from
// sum(weights) => len(numbers), so instead ranges are stored and binary search should be used at the pickIndex step.
type Solution struct {
	aggregateWeights [][]int
	totalWeight      int
}

// Time: O(n)
// Space: O(n)
func Constructor(weights []int) Solution {
	s := Solution{aggregateWeights: make([][]int, len(weights)), totalWeight: 0}
	for i, weight := range weights {
		s.aggregateWeights[i] = []int{s.totalWeight, s.totalWeight + weight - 1}
		s.totalWeight += weight
	}

	return s
}

// Time: O(log n) where n is len(weights)
// Space: O(1)
func (this *Solution) PickIndex() int {
	target := rand.Intn(this.totalWeight)
	return binarySearch(target, this.aggregateWeights, 0, len(this.aggregateWeights)-1)
}

func binarySearch(num int, aggregateWeights [][]int, min, max int) int {
	candidate := (min + max) / 2
	if num < aggregateWeights[candidate][0] {
		return binarySearch(num, aggregateWeights, min, candidate-1)
	}
	if num > aggregateWeights[candidate][1] {
		return binarySearch(num, aggregateWeights, candidate+1, max)
	}
	return candidate
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

func main() {
	obj := Constructor([]int{1, 3})

	targetToProbability := map[int]float64{
		0: 0.25,
		1: 0.75,
	}
	attempts := 1000000
	results := map[int]int{}
	for i := 0; i < attempts; i++ {
		idx := obj.PickIndex()
		results[idx]++
	}
	for idx, timesPicked := range results {
		actualProbability := float64(timesPicked) / float64(attempts)
		differenceToExpectedProbability := math.Abs(actualProbability - targetToProbability[idx])
		if differenceToExpectedProbability > 0.01 {
			fmt.Printf("Target %v was picked with %v probability, but %v was expected\n", idx, actualProbability, targetToProbability[idx])
		}
	}
}

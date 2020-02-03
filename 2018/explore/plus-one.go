package main

import (
	"fmt"
	"reflect"
)

// Time: O(n) Space: O(n)
func plusOne(digits []int) []int {
	if len(digits) == 1 && digits[0] == 0 {
		return []int{1}
	}
	newDigits := []int{}
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		newDigits = append(newDigits, sum%10)
		carry = sum / 10
	}
	if carry > 0 { // Edge case: 999; need to add a leading 1
		newDigits = append(newDigits, carry)
	}
	reverse(newDigits) // Alternatively use the original digits slice and add to the beginning
	return newDigits
}

// Time: O(n) Space: O(1)
func reverse(ns []int) {
	var (
		i = 0
		j = len(ns) - 1
	)
	for i < j {
		ns[i], ns[j] = ns[j], ns[i]
		i++
		j--
	}
}

func main() {
	fmt.Println(reflect.DeepEqual(plusOne([]int{9, 9, 9}), []int{1, 0, 0, 0}))
	fmt.Println(reflect.DeepEqual(plusOne([]int{1, 2, 3}), []int{1, 2, 4}))
	fmt.Println(reflect.DeepEqual(plusOne([]int{1, 2, 9}), []int{1, 3, 0}))
}

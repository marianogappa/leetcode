package main

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i, n := range nums {
		m[n] = i
	}

	for i, n := range nums {
		if v, ok := m[target-n]; ok && v != i {
			return []int{i, v}
		}
	}
	return nil // otherwise it doesn't compile
}

func main() {

}

package main

import "fmt"

// Pretty straightforward exercise, but lots of edge cases.
//
// It would be a tremendous waste of space to make every index have as many values as snaps, because one could call
// snap a million times without setting any values and use up all the memory. So the data structure for each index
// must be sparse. In order to make efficient sets & gets on sparse datasets, binarySearch is the go-to algorithm.
type SnapshotArray struct {
	arr  [][][]int
	snap int
}

// Time: O(1)
// Space: O(l)
func Constructor(length int) SnapshotArray {
	return SnapshotArray{arr: make([][][]int, length)}
}

// Time: O(log l)
// Space: O(1)
func (this *SnapshotArray) Set(index int, val int) {
	// If there's no set value at index, set the first value and return
	if len(this.arr[index]) == 0 {
		this.arr[index] = [][]int{{this.snap, val}}
		return
	}
	// If there are values at index, find the current snap value using binarySearch
	slot := binarySearch(this.arr[index], this.snap)
	// If binarySearch didn't find it, add a new value for current snap
	if slot[0] != this.snap {
		this.arr[index] = append(this.arr[index], []int{this.snap, val})
		return
	}
	// Otherwise, replace value for current snap
	slot[1] = val
}

// Time: O(1)
// Space: O(1)
func (this *SnapshotArray) Snap() int {
	this.snap++
	return this.snap - 1
}

// Time: O(log l)
// Space: O(1)
func (this *SnapshotArray) Get(index int, snap_id int) int {
	// Maybe there are no values for this index yet
	if len(this.arr[index]) == 0 {
		return 0
	}
	// Maybe the first recorded value for this index is for a greater snapId than requested
	if this.arr[index][0][0] > snap_id {
		return 0
	}
	// Otherwise binary search. It's ok to return a larger snapId value if snapId doesn't exist, but not a lower one.
	return binarySearch(this.arr[index], snap_id)[1]
}

// Time: O(log l)
// Space: O(1)
func binarySearch(arr [][]int, snapId int) []int {
	if len(arr) == 1 {
		return arr[0]
	}
	mid := len(arr) / 2
	if arr[mid][0] == snapId {
		return arr[mid]
	}
	if arr[mid][0] > snapId {
		return binarySearch(arr[:mid], snapId)
	}
	// There's a special case where mid snap is < but could still be the correct one if next snap is higher.
	// If we don't have a clause for this case, it will loop forever.
	if mid+1 < len(arr) && arr[mid+1][0] > snapId {
		return arr[mid]
	}
	return binarySearch(arr[mid:], snapId)
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
func main() {
	obj := Constructor(3)
	snapshotArr := &obj
	snapshotArr.Set(0, 5)
	fmt.Println(snapshotArr.Snap() == 0) // Take a snapshot, return snap_id = 0
	snapshotArr.Set(0, 6)
	fmt.Println(snapshotArr.Get(0, 0) == 5)

	obj = Constructor(1)
	snapshotArr = &obj
	snapshotArr.Set(0, 4)
	snapshotArr.Set(0, 16)
	snapshotArr.Set(0, 13)
	fmt.Println(snapshotArr.Snap() == 0) // Take a snapshot, return snap_id = 0
	snapshotArr.Get(0, 0)
	fmt.Println(snapshotArr.Snap() == 1) // Take a snapshot, return snap_id = 1

	obj = Constructor(2)
	snapshotArr = &obj
	fmt.Println(snapshotArr.Snap() == 0) // Take a snapshot, return snap_id = 0
	fmt.Println(snapshotArr.Get(1, 0) == 0)
	fmt.Println(snapshotArr.Get(0, 0) == 0)
	snapshotArr.Set(1, 8)
	fmt.Println(snapshotArr.Get(1, 0) == 0)
	snapshotArr.Set(0, 20)
	fmt.Println(snapshotArr.Get(0, 0) == 0)
	snapshotArr.Set(0, 7)
}

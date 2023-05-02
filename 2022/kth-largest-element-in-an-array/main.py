# Getting the kth largest/smallest item in O(n) is what quickselect does.
#
# Note that to get largest instead of the more usual smallest, you only
# need to change the "<" to ">" in the partition function. The comparisons
# in the quickselect function don't involve numbers, but indices, so they
# don't need to change.

from random import randint

# Time: O(n) average, O(n^2) worst
# Space: O(1) in-place
class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        return quickselect(nums, 0, len(nums)-1, k)
    
def partition(nums: List[int], left, right: int) -> int:
    if len == 1:
        return nums[0]

    pivot = randint(left, right)
    
    nums[pivot], nums[right] = nums[right], nums[pivot]
    
    for i in range(left, right):
        if nums[i] > nums[right]:
            nums[i], nums[left] = nums[left], nums[i]
            left += 1
    
    nums[right], nums[left] = nums[left], nums[right]
    
    return left
    
def quickselect(nums: List[int], left, right, k: int) -> int:
    if left == right:
        return nums[left]
    
    pivot = partition(nums, left, right)
    
    if pivot == k - 1:
        return nums[pivot]
    
    if pivot > k - 1:
        return quickselect(nums, left, pivot - 1, k)
    
    return quickselect(nums, pivot + 1, right, k)

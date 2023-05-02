# Intuition:
# 
# - If the digits are descending, there's no greater possible: [3, 2, 1]
# - If the digits are ascending, flipping the last two makes immediate greater: [1, 2, 3] -> [1, 3, 2]
# 
# Consider this:
# [4, 3, 2, 1, 4, 3, 2, 1] which should go to: [4, 3, 2, 2, 1, 1, 3, 4]
# 
# - Looks like first we look right to left for the first decreasing digit.
# - Once found, it must be swapped with a larger one on the right, but should be the "smallest larger".
# - The smallest larger MUST be the first larger going again right to left.
# - Once done, the numbers to the right of the swapped pivot have to be sorted ascendingly.
# - Because numbers are already sorted descendingly (even with the swap), rather than sort, just reverse!
class Solution:
    """

    
    Time: O(n)
    Space: O(1)
    
    """
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        idx = find_idx_first_descending_right_to_left(nums)
        if idx == -1:
            reverse_from(nums, 0)
            return
        
        larger_idx = find_idx_first_larger_right_to_left(nums, nums[idx])
        
        nums[idx], nums[larger_idx] = nums[larger_idx], nums[idx]
        
        reverse_from(nums, idx+1)

def find_idx_first_descending_right_to_left(nums: list[int]) -> int:
    for i in range(len(nums)-2, -1, -1):
        if nums[i] < nums[i+1]:
            return i
    return -1

def reverse_from(nums: list[int], left: int) -> None:
    right = len(nums)-1
    while left < right:
        nums[left], nums[right] = nums[right], nums[left]
        left += 1
        right -= 1

def find_idx_first_larger_right_to_left(nums: list[int], num: int) -> int:
    for i in range(len(nums)-1, -1, -1):
        if nums[i] > num:
            return i
    # Unreachable code
    return -1

# A trivial constant space two pass is to count zeroes, ones & twos and overwrite on a second pass.
#
# Here's how to solve it in one pass: (remember the pivot on quicksort)
# - Use 3 pointers: `low` & `high` on either side, and `mid` going through the list
# - As `mid` moves, swap 0s & 2s with the `low` & `high` pivots and update them
# - BEWARE: When `low` advances, it pushes `mid` too!

class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        low = 0
        mid = 0
        high = len(nums)-1

        while mid <= high:
            # If it's a 0, swap with low pivot and increment both
            if nums[mid] == 0:
                nums[mid], nums[low] = nums[low], nums[mid]
                low += 1
                # Why increment mid only in this case but not when high moves?
                # mid must always be after low, so low pushes it when it moves
                mid += 1 
            
            # If it's a 2, swap with high pivot and decrement high
            elif nums[mid] == 2:
                nums[mid], nums[high] = nums[high], nums[mid]
                high -= 1
            
            # If it's a 1, leave it and increment mid
            else:
                mid += 1

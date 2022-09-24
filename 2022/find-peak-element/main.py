# Time: O(log(n))
# Space: O(log(n)) can be O(1) on iterative binary search
#
# This constraint is providing an elusive hint: nums[i] != nums[i + 1] for all valid i
# Until you figure out why they mention it, you won't find the solution.
#
# Intuition: given any pair of numbers, one will be larger than the other. At that point
# you know there will be a local peak in that direction (i.e. either when the ascending
# trend reverses, or when you escape the array). With this knowledge, you can do
# binary search.
class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        return binary_search(nums, 0, len(nums)-1)

def binary_search(nums: list[int], left, right: int) -> Optional[int]:
    if left == right:
        return left
    
    # A trick I learned on this exercise about binary search:
    # Since we know right-left >= 1, we can guarantee mid+1 exists!
    # So whenever binary search requires comparing to a subsequent element,
    # choose the element on the right!
    
    mid = (left + right) // 2
    if nums[mid] > nums[mid+1]:
        # Note that mid CAN be the solution, so don't eliminate it.
        return binary_search(nums, left, mid)
    elif nums[mid] < nums[mid+1]:
        return binary_search(nums, mid+1, right)
    
from typing import List

# Time: O(log(n) + k)
# Space: O(log(n)) or O(k) if solution space counts; log(n) stack space removed with iterative Binary Search
class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        # Find minimal distance with binary search
        mid_idx = binary_search(arr, x, 0, len(arr)-1)
        # Advance to the left & right with 2 pointers to "enlarge a range" until range has k elements
        left, right = find_range(arr, mid_idx, k, x)
        # Return a slice of the array as defined by the range
        return arr[left:right+1]
        
def binary_search(arr: list[int], x, left, right: int) -> int:
    if left == right:
        return left
    
    # Because in neither recursion case we remove mid, we might end up with len == 2 forever
    if right == left + 1:
        return right if abs(arr[left]-x) > abs(arr[right]-x) else left

    mid = (left + right) // 2
    diff = arr[mid]-x
    
    if diff == 0:
        return mid
    
    if diff > 0:
        return binary_search(arr, x, left, mid)
    return binary_search(arr, x, mid, right)

def find_range(arr: list[int], mid_idx: int, k, x: int) -> tuple[int, int]:
    left = right = mid_idx
    while right-left+1 < k:
        if left == 0:
            right += 1
        elif right == len(arr) - 1:
            left -= 1
        elif abs(arr[left-1] - x) <= abs(arr[right+1] - x):
            left -= 1
        else:
            right += 1

    return (left, right)

print(Solution().findClosestElements([1,2,3,4,5], 4, 3), " == [1, 2, 3, 4]")
print(Solution().findClosestElements([1,2,3,4,5], 4, -1), " == [1, 2, 3, 4]")
print(Solution().findClosestElements([1,1,1,10,10,10], 1, 9), " == [10]")

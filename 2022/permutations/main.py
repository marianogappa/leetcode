# No point describing permutations. Just read the code. It's a DFS with burning and backtracking.

class Solution:
    # Time complexity: O(n!) because we start with `n` options, then `n-1`, then `n-2`, etc.
    # Space complexity: O(n!) if solution space counts, O(n) if not.
    def permute(self, nums: List[int]) -> List[List[int]]:
        return do_permute(nums, [], [])

BURNED = float("inf")

def do_permute(nums: List[int], partial: List[int], result: List[List[int]]) -> List[List[int]]:
    # A new permutation is reached when partial is the same length as nums
    if len(partial) == len(nums):
        result.append(partial.copy()) # Remember to clone! We'll mutate partial later.
        return result

    for i, num in enumerate(nums):
        if num == BURNED:
            continue
        
        # Pick this number and burn it
        partial.append(num)
        nums[i] = BURNED

        do_permute(nums, partial, result)
        
        # Don't forget to unburn the number and backtrack 
        nums[i] = num
        partial.pop()

    return result

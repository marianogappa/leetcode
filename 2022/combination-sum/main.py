# Trivial DFS except for two key points:
#
# - When recursing, only allow candidates at the current index or higher: this avoids duplicates
# - Reuse the partial solution by backtracking, and clone when a solution is found.

class Solution:
    # Time: O(n^(m/t)) n == len(candidates), m == smallest candidate, t == target
    # Space: O(m/t) being the len of the longest combination (solution space not counting as space)
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        dfs(candidates, target, [], result)
        return result
    
def dfs(candidates: List[int], target: int, current: list[int], result: list[list[int]]) -> list[list[int]]:
    if target < 0:
        return
    if target == 0:
        result.append(current.copy())
        return
    
    for i, candidate in enumerate(candidates):
        current.append(candidate)
        dfs(candidates[i:], target-candidate, current, result)
        current.pop()

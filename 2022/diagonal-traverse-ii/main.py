from collections import defaultdict
from typing import List

# You pretty much have to know the trick to pull this one off, but then it's easy.
# The problem with the intuitive strategy is that the data structure is sparse, and you cannot
# afford to check all empty slots or it will TLE.
#
# The trick is that each diagonal line of cells has the same unique row+col sum, so as long as those
# cells are traversed in order to each other, it is not necessary to traverse the lists in diagonal
# order. You can achieve this by traversing the lists bottom-up, but each list can be traversed
# normally: left to right. Use a hashmap to keep per-sum lists. Each list will have the
# diagonally-traversed set of cells of each sum.
#
# Time: O(n)
# Space: O(n)
class Solution:
    def findDiagonalOrder(self, nums: List[List[int]]) -> List[int]:
        x_y_sum_to_list: dict[int, list[int]] = defaultdict(list[int])
        max_sum = float("-inf")
        for y in range(len(nums)-1, -1, -1):
            max_sum = max(max_sum, len(nums[y])-1+y)
            for x in range(len(nums[y])):
                x_y_sum_to_list[x+y].append(nums[y][x])

        result = []
        for cur_sum in range(0, max_sum+1):
            result += x_y_sum_to_list[cur_sum]

        return result

print(Solution().findDiagonalOrder([[1,2,3],[4,5,6],[7,8,9]]) == [1,4,2,7,5,3,8,6,9])
print(Solution().findDiagonalOrder([[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]) == [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16])

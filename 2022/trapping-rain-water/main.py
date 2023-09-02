# The reason a linear solution is unintuitive is that, as you loop through the array, you need to know the maximum height
# to the left AND to the right, to know how much the current position can trap. How to know without backtracking?
# 
# If it was possible, it requires a two-pointer solution.
# 
# The trick:
# - As you move, keep maximums on both sides: left_max & right_max.
# - EUREKA: Only move the pointer that has the smaller maximum 🤯!
# - This way, you know how much water the current position can trap: `min(left_max, right_max) - height[current]`

# Time: O(n)
# Space: O(1)
class Solution:
    def trap(self, height: List[int]) -> int:
        left = 0
        right = len(height) - 1
        trapped = 0
        left_max = 0
        right_max = 0

        while left < right:
            left_max = max(left_max, height[left])
            right_max = max(right_max, height[right])

            if left_max < right_max:
                trapped += min(left_max, right_max) - height[left]
                left +=1
            else:
                trapped += min(left_max, right_max) - height[right]
                right -= 1

        return trapped

# A two-pointer solution is intuitive, but the question is how to not lose any candidates.
#
# Start with two pointers on either side, and advance the pointer with the lesser height.

class Solution:
    # Time: O(n)
    # Space: O(1)
    def maxArea(self, height: List[int]) -> int:
        max_area = float("-inf")

        # Start with two pointers on either side
        left = 0
        right = len(height) -1

        while left < right:
            # Keep track of the max area seen
            max_area = max(max_area, (right - left) * min(height[left], height[right]))

            # Advance the pointer with the lesser height
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        
        return max_area

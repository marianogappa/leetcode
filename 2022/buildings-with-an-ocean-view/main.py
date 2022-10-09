# Time: O(n)
# Space: O(n)
class Solution:
    def findBuildings(self, heights: List[int]) -> List[int]:
        result = [len(heights)-1]
        tallest = heights[-1]

        for i in range(len(heights)-2, -1, -1):
            if tallest < heights[i]:
                result.append(i)
                tallest = heights[i]
        return result[::-1]

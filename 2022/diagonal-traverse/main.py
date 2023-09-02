# wip!
# Since it's tricky to iterate in the required pattern, it's easier to always iterate
# with the same delta (up-right), and reverse every second sub-path (linear time).
#
# At that point, it's trivial: traverse first column and last row always towards up-right.

# Time: O(n)
# Space: O(n)
class Solution:
    def findDiagonalOrder(self, mat: List[List[int]]) -> List[int]:
        path: list[int] = []
        
        should_reverse = False
        for y in range(len(mat)):
            path += traverse(mat, [0, y], [1, -1], should_reverse)
            should_reverse = not should_reverse
        
        for x in range(1, len(mat[y])):
            path += traverse(mat, [x, y], [1, -1], should_reverse)
            should_reverse = not should_reverse
        
        return path
        
def traverse(mat: list[list[int]], pos: list[int], delta: list[int], should_reverse: bool) -> list[int]:
    path: list[list[int]] = []
    while is_in_bounds(mat, pos):
        path.append(mat[pos[1]][pos[0]])
        pos[0] += delta[0]
        pos[1] += delta[1]
    
    return reverse(path) if should_reverse else path
        
def is_in_bounds(mat: list[list[int]], pos: list[int]) -> bool:
    return pos[0] >= 0 and pos[1] >= 0 and pos[1] < len(mat) and pos[0] < len(mat[pos[1]])

def reverse(array: list[int]) -> list[int]:
    left = 0
    right = len(array)-1
    
    while left < right:
        array[left], array[right] = array[right], array[left]
        left += 1
        right -= 1
    
    return array

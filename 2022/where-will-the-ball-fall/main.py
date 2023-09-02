# wip!
# Intuition #1: movement algorithm is ALWAYS: moving once horizontally (left or right) and then once vertically (down).
# Intuition #2: the delta of horizontal movement is given by the number in the cell: next_x = x + grid[y][x].
# Intuition #3: you can only move horizontally if the source and destination cells have the same number!
from typing import List

# Time: O(x*y) solve for every x, each solution iterates y times
# Space: O(1) or O(x) if solution space counts
class Solution:
    def findBall(self, grid: List[List[int]]) -> List[int]:
        # Solve exercise for every starting column
        return [solve_from(grid, x) for x in range(len(grid[0]))]

def solve_from(grid: List[List[int]], x: int) -> int:
    # For every row...
    for y in range(len(grid)):
        # First tentatively move horizontally. Direction is given by cell's value (1 => right, -1 => left) as a delta.
        next_x = x + grid[y][x]

        # If by moving we went past a wall, or if the column we moved to has a different value, done (i.e. we're stuck).
        if not is_within_walls(grid, next_x) or not can_move_horizontally(grid, y, x, next_x):
            return -1
        
        # Otherwise, we CAN move horizontally. Update x, and we can ALWAYS go down. So iterate (i.e. move down).
        x = next_x

    # If we didn't hit a "return -1" then we reached the end. The x is the column we ended up in.
    return x

def is_within_walls(grid: List[List[int]], x: int) -> bool:
    return x >= 0 and x < len(grid[0])

def can_move_horizontally(grid: List[List[int]], y, x, next_x: int):
    return grid[y][next_x] == grid[y][x]

print(Solution().findBall([[1,1,1,-1,-1],[1,1,1,-1,-1],[-1,-1,-1,1,1],[1,1,1,1,-1],[-1,-1,-1,-1,-1]]))
print(Solution().findBall([[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1],[1,1,1,1,1,1],[-1,-1,-1,-1,-1,-1]]))

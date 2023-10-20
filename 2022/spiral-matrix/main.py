# It's possible to solve in O(1) space but this is more readable.

class Solution:
    # Time: O(n)
    # Space: O(n) for the seen set
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        result = []

        # A set of visited cells
        seen = set([])

        # Possible directions (with an index)
        deltas = [(1, 0), (0, 1), (-1, 0), (0, -1)]
        di = 0
        
        # Cursor
        x = 0
        y = 0
        
        total_len = len(matrix)*len(matrix[0])
        while len(result) < total_len:
            # While out of bounds or in seen cells
            while (x, y) in seen or x < 0 or y < 0 or y >= len(matrix) or x >= len(matrix[0]):
                # Backtrack
                x -= dx
                y -= dy
                
                # Change direction
                di = (di + 1) % 4

                # Advance
                dx, dy = deltas[di]
                x += dx
                y += dy

            # Otherwise advance in the current direction
            seen.add((x, y))
            result.append(matrix[y][x])
            dx, dy = deltas[di]
            x += dx
            y += dy
        
        return result

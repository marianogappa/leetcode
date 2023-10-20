# The rottenness spreads to adjacent oranges minute by minute. This sounds like BFS!
#
# - Root nodes (level 0) are the rotten oranges.
# - Each node connects to all adjacent oranges on the NEXT LEVEL (or minute).
# - Make sure to "burn" oranges as you append them to the next level, to not overcount or cycle.
# - At the end of BFS, check if there are oranges. If there are, return -1.

class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        # Start BFS with all rotten oranges in the queue
        queue = deque([(x, y) for y in range(len(grid)) for x in range(len(grid[0])) if grid[y][x] == 2])
        level = 0
        while queue:
            for _ in range(len(queue)):
                x, y = queue.popleft()

                for neighbor_delta in [(-1, 0), (0, -1), (1, 0), (0, 1)]:
                    dx, dy = neighbor_delta
                    nx, ny = x + dx, y + dy

                    # Ignore out of bounds!
                    if (
                        nx < 0 or ny < 0 or 
                        ny >= len(grid) or nx >= len(grid[0]) or grid[ny][nx] != 1
                    ):
                        continue

                    grid[ny][nx] = 2
                    queue.append((nx, ny))
            # This level counts only if we burned oranges. If we did, queue is not empty!
            if queue:
                level += 1

        # Don't forget to check if there are isolated oranges!
        if len([True for y in range(len(grid)) for x in range(len(grid[0])) if grid[y][x] == 1]):
            return -1

        return level
            


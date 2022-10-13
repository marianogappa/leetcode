class Solution:
    """
    Time: O(mnk)
    Space: O(mn)

    BFS works better, because on each level "steps += 1", and when you reach target it HAS to be
    the shortest way.

    The only trick is how to deal with obstacles. Obstacles are the reason Dijkstra & DFS are
    problematic.

    In BFS & DFS, there's the issue of arriving at the same node twice, so a "visited" set is
    necessary. However, one may go on one path and reach two many obstacles, and then try a
    different path that is longer but has less obstacles and reaches target! So how to solve it?

    Make the "visited" set a dict, and store how many obstacles can be eliminated. Instead of not
    allowing visiting a visited node, allow it only if the current number of obstacles left is
    higher than the last time it was visited.
    """
    def shortestPath(self, grid: List[List[int]], k: int) -> int:
        queue = deque()
        queue.append((0, 0, 0, k))
        
        deltas = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        visited = {}
        target = (len(grid[0])-1, len(grid)-1)
        
        while queue:
            for _ in range(len(queue)):
                x, y, steps, obstacles_left = queue.popleft()

                if (x, y) == target:
                    return steps

                if (x, y) in visited and obstacles_left <= visited[(x, y)]:
                    continue
                                
                visited[(x, y)] = obstacles_left

                for delta in deltas:
                    new_x, new_y = x + delta[0], y + delta[1]

                    if not is_in_bounds(new_x, new_y, grid):
                        continue

                    if not obstacles_left and grid[new_y][new_x] == 1:
                        continue

                    queue.append((new_x, new_y, steps + 1, obstacles_left - int(grid[new_y][new_x] == 1)))
        
        return visited.get(target, -1)
            
def is_in_bounds(x: int, y: int, grid: List[List[int]]) -> bool:
    return x >= 0 and y >= 0 and y < len(grid) and x < len(grid[0])

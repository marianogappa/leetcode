# Take the hint that distance to a zero in an interconnected grid is a graph problem!
#
# Since each edge (level!) increases distance by 1, BFS is the most intuitive solution!
#
# - Level 0 of BFS are all nodes with zeroes. 
# - Each UNSEEN neighbor will have distance 1, and so on.
# - Simply do BFS with a `seen` set to avoid cycles.

class Solution:
    # Time: O(n) because all nodes will be processed once
    # Space: O(n) because seen & queue will contain all nodes
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        max_y = len(mat)
        max_x = len(mat[0])
        zeroes = [(x, y) for y in range(max_y) for x in range(max_x) if mat[y][x] == 0]
        queue = deque(zeroes)
        seen = set([])
        level = -1

        while queue:
            level += 1
            # Only iterate through current level (nodes will be added below!!)
            for _ in range(len(queue)):
                x, y = queue.popleft()

                # Only pursue unseen nodes
                if (x, y) in seen:
                    continue
                seen.add((x, y))

                mat[y][x] = level

                for neighbor_delta in [(-1, 0), (0, -1), (1, 0), (0, 1)]:
                    dx, dy = neighbor_delta
                    nx, ny = (x + dx, y + dy)

                    # Ignore out of bounds
                    if nx < 0 or ny < 0 or nx >= max_x or ny >= max_y:
                        continue
                    
                    # Add neighbor to next level!
                    queue.append((nx, ny))
        
        return mat

# Since the grid is an undirected graph and we must find how many connected groups there are,
# this is a textbook case for UnionFind!
#
# - Add all lands
# - Union all connections 
# - Count how many groups there are!

class Solution:
    # Time: O(n) because we iterate over the grid twice doing constant ops, and then O(n) set_count.
    # Space: O(n) because the union-find DS requires O(n) space.
    def numIslands(self, grid: List[List[str]]) -> int:
        uf = UnionFind()
        
        # Add all lands to the union-find DS.
        for y in range(len(grid)):
            for x in range(len(grid[0])):
                if grid[y][x] != "1":
                    continue
                uf.add((x, y))
        
        # Add all edges (connections) between lands to the union-find DS.
        # Note that we only need to check right & down. Can also connect left & up, but redundant.
        for y in range(len(grid)):
            for x in range(len(grid[0])):
                if grid[y][x] != "1":
                    continue
                for delta in [(1, 0), (0, 1)]:
                    dx, dy = delta
                    nx, ny = (x+dx, y+dy)
                    # Don't go out of bounds!
                    if ny >= len(grid) or nx >= len(grid[0]):
                        continue
                    if grid[ny][nx] != "1":
                        continue
                    uf.union((x, y), (nx, ny))
        
        # The number of connected groups is the number of islands!
        return uf.set_count()

class UnionFind: # Used to keep track of disjointed sets on graphs.
    parent: list[int] # Sets are trees. Each vertex has a parent.
    size: list[int] # Optimization: on union, larger set becomes parent.
    idx: dict[any, int] # Optional: elements are mapped to int "labels".

    def __init__(self):
        self.parent = []
        self.size = []
        self.idx = {}

    # Add a new vertex to the forest, as a new tree of size 1
    def add(self, vertex: any):
        idx = len(self.parent) # Next available idx to use for this vertex.
        self.idx[vertex] = idx
        self.parent.append(idx) # New vertex is its own parent in own tree.
        self.size.append(1) # Tree of one vertex: size = 1

    # Union two vertices: if they belong to != sets, make larger parent
    # of the smaller. Keep track of sizes.
    # vertices must exist (via add)!
    def union(self, vertex1: any, vertex2: any):
        set1 = self.find(vertex1)
        set2 = self.find(vertex2)
        if set1 != set2:
            if self.size[set1] > self.size[set2]:
                self.size[set1] += self.size[set2]
                self.parent[set2] = set1
            else:
                self.size[set2] += self.size[set1]
                self.parent[set1] = set2

    def _find(self, idx: int) -> int:
        return self._find(self.parent[idx]) if self.parent[idx] != idx else idx
    
    # Find the set a vertex belongs to.
    # vertex must exist (via add)!
    def find(self, vertex: any) -> int:
        return self._find(self.idx[vertex])
    
    # Count how many different sets exist (O(n))
    def set_count(self) -> int:
        return sum(1 for idx, parent in enumerate(self.parent) if idx == parent)

# If you have trouble remembering how to implement UnionFind, there's a simple DFS solution too:
#
# - For every cell, if it's land, +1 to the total number of islands, and then BURN connected islands.
# - BURN: dfs via adjacent cells and replace the "1" to "X" or whatever.
# - Note that the "burning" plays the role of the "visited" set, so you don't stack overflow.

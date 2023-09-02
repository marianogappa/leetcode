# wip!
# If any two stones share a col or row, they're part of the same "group".
# In a group, you can always iteratively remove stones in some order until there's only one left.
#
# Therefore:
# 1) Use union-find to answer how many groups there are
# 2) Number of stones that can be removed: total stones - total groups
#
# Trick: the "1s" in (1, 1) are different individuals, so add a large number to all "y"s to make them different.

# Time: O(nlog(n)) due to 2n operations on n nodes
# Space: O(n)
class Solution:
    def removeStones(self, stones: List[List[int]]) -> int:
        uf = UnionFind()
        for stone in stones:
            # Must add a large number to second coordinate to make sure numbers don't overlap
            uf.union(stone[0], 10000 + stone[1])
        
        group_count = len({uf.find(key) for key in uf.parent})
        
        return len(stones) - group_count
    
# This implementation of UnionFind uses a dict instead of an array but it
# has the advantage that it can grow arbitrarily.
class UnionFind:
    def __init__(self):
        self.parent: dict[int, int] = {}
    
    def union(self, a, b: int) -> None:
        # Initialise a & b into the dict only if they don't exist already
        self.parent.setdefault(a, a)
        self.parent.setdefault(b, b)
        
        # Union them in an arbitrary order
        self.parent[self.find(a)] = self.find(b)
    
    def find(self, a: int) -> int:
        # Recursively find parents until we reach the main parent
        return self.find(self.parent[a]) if self.parent[a] != a else a

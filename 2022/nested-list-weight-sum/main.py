# Time: O(n)
# Space: O(d) depth
class Solution:
    def depthSum(self, nestedList: List[NestedInteger]) -> int:
        return sum([ flatten(item, 1) for item in nestedList ])

def flatten(item: NestedInteger, depth: int) -> int:
    if item.isInteger():
        return item.getInteger() * depth
    return sum([flatten(nested_item, depth+1) for nested_item in item.getList()])

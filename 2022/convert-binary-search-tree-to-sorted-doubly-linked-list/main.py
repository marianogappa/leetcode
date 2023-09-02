# wip!
class Solution:
    def treeToDoublyList(self, root: 'Optional[Node]') -> 'Optional[Node]':
        return array_to_sdll(bst_to_array(root, []))

def bst_to_array(root: 'Optional[Node]', partial: list[int]) -> list['Optional[Node]']:
    if not root:
        return partial
    
    bst_to_array(root.left, partial)
    partial.append(root)
    bst_to_array(root.right, partial)

    return partial

def array_to_sdll(nums: list['Optional[Node]']) -> 'Optional[Node]':
    if not nums:
        return None
    
    if len(nums) == 1:
        aux = nums[0]
        aux.left = aux.right = aux
        return aux
    
    for i in range(0, len(nums)-1):
        nums[i].left = nums[i-1]
        nums[i].right = nums[i+1]
    
    nums[-1].right = nums[0]
    nums[-1].left = nums[-2]
    
    return nums[0]

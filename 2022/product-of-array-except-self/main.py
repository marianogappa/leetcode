# Strategy is to do a "running product" in the `range(start+1, end+1)`
# and then in the range `(end-1, start-1, -1)`, and multiply those together.
#
# For example `[1, 2, 3, 4]``:
#
# - start with the zero value of multiplication: `[1, 1, 1, 1]`
# - run a running product from `(start+1 -> end)`: `[1, 1*1, 1*1*2, 1*1*2*3] = [1, 1, 2, 6]`
# - run a running product from `(end-1 -> start)`: `[1*4*3*2, 1*4*3, 1*4, 1] = [24, 12, 4, 1]`
# - multiply those together: `[1*24, 1*12, 2*4, 6*1] = [24, 12, 8, 6]`

# Time: O(n)
# Space: O(1) Note: exercise says "The output array does not count as extra
#                   space for space complexity analysis."
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        product = [1] * len(nums)
        
        running_product = 1
        for i in range(1, len(nums)):
            running_product *= nums[i-1]
            product[i] *= running_product
        
        running_product = 1
        for i in range(len(nums)-2, -1, -1):
            running_product *= nums[i+1]
            product[i] *= running_product

        return product

# wip!
# Whenever an exercise says "kth element", always think of heaps.
#
# Put the first row in a min-heap. Then, "replace" k times (i.e. pop and push immediately) the min element with the element
# right below it. Both operations should be logarithmic (to the size of the heap which is one row). The k-th value popped
# is the one to return.
#
# TODO: investigate the constant memory solution of using binary search.

import heapq

# Time: O(k log r) where r is the length of a row of the matrix
# Space: O(r)
class Solution:
    def kthSmallest(self, matrix: list[list[int]], k: int) -> int:
        # Space: O(r) where r is the size of a row. We will replace items but won't add any so that's the space complexity.
        heap = []
        for i in range(len(matrix[0])):
            # Make the heap element a tuple containing the value but also the (y, x), so that when we pop, we can replace
            # with the element below it.
            heap.append((matrix[0][i], 0, i))

        # Time: O(r) heapify is linear time to size of heap which is one row at the moment
        heapq.heapify(heap)

        # Time: O(k log r) replace the min-element in the heap with the one below it
        # Space: O(1) no new space added, just replaced
        val = 0
        for i in range(k):
            (val, y, x) = heapq.heappop(heap)
            if y+1 < len(matrix):
                heapq.heappush(heap, (matrix[y+1][x], y+1, x))

        return val


print(Solution().kthSmallest([[-5]], 1) == -5)
print(Solution().kthSmallest([[1, 5, 9], [10, 11, 13], [12, 13, 15]], 8) == 13)
print(Solution().kthSmallest([[1, 3, 5], [6, 7, 12], [11, 14, 14]], 6) == 11)

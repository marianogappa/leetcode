# wip!
# Time: O(n)
# Space: O(1)
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        length = len(matrix)

        if length <= 1:
            return

        for row in matrix:
            self.reverse_row(row)

        # Draw a mental line on the diagonal starting bottom-left to top-right, and mirror numbers to both sides of it.
        for y in range(length):
            for x in range(length):
                # Only need to do mirroring operations to the "left" of the diagonal.
                if x + y >= length - 1:
                    break
                # Flip mirroring numbers
                matrix[y][x], matrix[length-1-x][length-1 -
                                                 y] = matrix[length-1-x][length-1-y], matrix[y][x]

    def reverse_row(self, row: List[int]) -> None:
        start = 0
        end = len(row)-1
        while start < end:
            row[start], row[end] = row[end], row[start]
            start += 1
            end -= 1

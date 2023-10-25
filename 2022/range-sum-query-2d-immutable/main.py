class NumMatrix:

    def __init__(self, matrix: List[List[int]]):
        self.s = [[0 for _ in range(len(matrix[0])+1)] for _ in range(len(matrix)+1)]
        for y in range(len(matrix)-1, -1, -1):
            for x in range(len(matrix[0])-1, -1, -1):
                self.s[y][x] = (
                    matrix[y][x] +
                    self.s[y+1][x] + 
                    self.s[y][x+1] -
                    self.s[y+1][x+1]
                )

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        return self.s[row1][col1] + self.s[row2+1][col2+1] - self.s[row1][col2+1] - self.s[row2+1][col1]


# Your NumMatrix object will be instantiated and called as such:
# obj = NumMatrix(matrix)
# param_1 = obj.sumRegion(row1,col1,row2,col2)

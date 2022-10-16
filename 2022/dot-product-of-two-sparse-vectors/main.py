# Time: O(n) 
# Space: O(n)
class SparseVector:
    def _init_(self, nums: List[int]):
        self.idx_to_num = []
        for i in range(len(nums)):
            if (nums[i] != 0):
                self.idx_to_num.append((i, nums[i]))
    
    # Return the dotProduct of two sparse vectors
    def dotProduct(self, vec: 'SparseVector') -> int:
        result = 0
        i, j = 0, 0
        while (i < len(self.idx_to_num) and j < len(vec.idx_to_num)) :
            if(self.idx_to_num[i][0] < vec.idx_to_num[j][0]):
                i += 1
            elif(self.idx_to_num[i][0] > vec.idx_to_num[j][0]):
                j += 1
            else:
                result += self.idx_to_num[i][1] * vec.idx_to_num[j][1]
                i += 1
                j += 1
        return result

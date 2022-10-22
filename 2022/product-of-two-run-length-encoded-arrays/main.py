from typing import List

# Time: O(n) where n == len(encoded1)
# Space: O(n)
#
# Decompressing the RLEs is not feasible due to space constraints. Try to multiply the RLEs directly.
#
# Traverse the encoded lists with two pointers, and decompose the "current" tuple with the largest count into two
# tuples, such that the first one matches the count of the one with the smallest count.
class Solution:
    def findRLEArray(self, encoded1: List[List[int]], encoded2: List[List[int]]) -> List[List[int]]:
        return self.multiply(encoded1, encoded2)
    
    def multiply(self, encoded1: List[List[int]], encoded2: List[List[int]]) -> List[List[int]]:
        result = []
        i = j = 0
        while i < len(encoded1):
            # Easy case: counts match, so just multiply
            if encoded1[i][1] == encoded2[j][1]:
                self.result_append(result, [encoded1[i][0] * encoded2[j][0], encoded1[i][1]])
                i += 1
                j += 1
            # Counts don't match: split larger count into two tuples, "append" the first one and advance the pointer
            # of the smallest tuple.
            else:
                self.result_append(result, [encoded1[i][0] * encoded2[j][0], min(encoded1[i][1], encoded2[j][1])])
                if encoded1[i][1] > encoded2[j][1]:
                    encoded1[i][1] -= encoded2[j][1]
                    j += 1 
                else:
                    encoded2[j][1] -= encoded1[i][1]
                    i += 1
        
        return result

    # Solution TLEs if we append all tuples and merge the subsequent ones with the same number later, so we must merge
    # as we append.
    def result_append(self, result:List[List[int]], elem:List[int]):
        if not result or result[-1][0] != elem[0]:
            result.append(elem)
        else:
            result[-1][1] += elem[1] 

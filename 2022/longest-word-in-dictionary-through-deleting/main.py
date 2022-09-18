from typing import List

# Time: O(d*(min(s,e))) where s => len(s), d => len(dictionary), e => len of largest word in dictionary
# Space: O(1)
class Solution:
    def findLongestWord(self, s: str, dictionary: List[str]) -> str:
        longest = ''
        for word in dictionary:
            if is_subset(word, s) and is_longer(word, longest):
                longest = word
        
        return longest

def is_longer(word, longest: str) -> str:
    return longest == '' or len(word) > len(longest) or (len(word) == len(longest) and word < longest)

def is_subset(needle, haystack: str) -> bool:
    i = 0
    for char in haystack:
        if i == len(needle):
            break
        if char == needle[i]:
            i += 1
    
    return i == len(needle)

print(Solution().findLongestWord('abpcplea', ["ale","apple","monkey","plea"]),  "== apple") 
print(Solution().findLongestWord('abpcplea', ["a","b","c"]) == 'a') 

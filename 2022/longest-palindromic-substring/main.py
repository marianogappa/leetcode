# There exists a Manacher algorithm that can solve this problem in O(n) time. It's not worth learning it.
#
# A readable option is to solve twice for odd and even length palindromes.
#
# You can loop through the center of the palindrome and expand outwards while palindromic, and keep a max length.

class Solution:
    # Time: O(n^2)
    # Space: O(1)
    def longestPalindrome(self, s: str) -> str:
        longest_odd = longest_odd_palindrome(s)
        longest_even = longest_even_palindrome(s)

        return longest_odd if len(longest_odd) > len(longest_even) else longest_even

def longest_odd_palindrome(s: str) -> str:
    longest = s[0]
    for start in range(len(s)):
        expand = 0
        # While not out of bounds and still palindromic, keep track of max & expand
        while start - expand >= 0 and start + expand < len(s) and s[start-expand] == s[start+expand]:
            if 2*expand + 1 > len(longest):
                longest = s[start - expand: start + expand + 1]
            expand += 1
    return longest

def longest_even_palindrome(s: str) -> str:
    longest = ""
    for start in range(len(s)-1):
        expand = 1
        # While not out of bounds and still palindromic, keep track of max & expand
        while start - expand + 1 >= 0 and start + expand < len(s) and s[start-expand+1] == s[start+expand]:
            if 2*expand > len(longest):
                longest = s[start - expand + 1: start + expand + 1]
            expand += 1
    return longest
